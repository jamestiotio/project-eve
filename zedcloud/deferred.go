// Copyright (c) 2018 Zededa, Inc.
// All rights reserved.

// Support for deferring sending of messages after a failure

package zedcloud

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"github.com/zededa/go-provision/flextimer"
	"time"
)

// Example usage:
// deferredChan := zedcloud.InitDeferred()
// select {
//      case change := <- deferredChan:
//		zedcloud.HandleDeferred(change)
// Before or after sending success call:
//	zedcloud.RemoveDeferred(key)
// After failure call
// 	zedcloud.SetDeferred(key, data, url, zedcloudCtx)
// or AddDeferred to build a queue for each key

type deferredItem struct {
	data        []byte
	url         string
	zedcloudCtx ZedCloudContext
	ignore400   bool
}

type deferredItemList struct {
	list []deferredItem
}

const longTime1 = time.Hour * 24
const longTime2 = time.Hour * 48

// Some day we might return this; right now only for the defaultCtx
type DeferredContext struct {
	deferredItems map[string]deferredItemList
	ticker        flextimer.FlexTickerHandle
}

// From first InitDeferred
var defaultCtx *DeferredContext

// Create and return a channel to the caller
func InitDeferred() <-chan time.Time {
	if defaultCtx != nil {
		log.Fatal("InitDeferred called twice")
	}
	defaultCtx = initImpl()
	return defaultCtx.ticker.C
}

func initImpl() *DeferredContext {
	ctx := new(DeferredContext)
	ctx.deferredItems = make(map[string]deferredItemList)
	// We always keep a flextimer running so that we can return
	// the associated channel. We adjust the times when we start and stop
	// the timer.
	ctx.ticker = flextimer.NewRangeTicker(longTime1, longTime2)
	return ctx
}

// Try to send all deferred items. Give up if any one fails
// Stop timer is map becomes empty
func HandleDeferred(event time.Time) {
	if defaultCtx == nil {
		log.Fatal("HandleDeferred no defaultCtx")
	}
	defaultCtx.handleDeferred(event)
}

func (ctx *DeferredContext) handleDeferred(event time.Time) {

	log.Infof("HandleDeferred(%v) map %d\n",
		event, len(ctx.deferredItems))
	iteration := 0 // Do some load spreading
	for key, l := range ctx.deferredItems {
		log.Infof("Trying to send for %s len %d\n", key, len(l.list))
		failed := false
		for i, item := range l.list {
			if item.data == nil {
				continue
			}
			log.Infof("Trying to send for %s item %d data len %d\n",
				key, i, len(item.data))
			resp, _, err := SendOnAllIntf(item.zedcloudCtx, item.url,
				int64(len(item.data)), bytes.NewBuffer(item.data),
				iteration, item.ignore400)
			if item.ignore400 && resp != nil &&
				resp.StatusCode >= 400 && resp.StatusCode < 500 {
				log.Infof("HandleDeferred: for %s ignore code %d\n",
					key, resp.StatusCode)
			} else if err != nil {
				log.Infof("HandleDeferred: for %s failed %s\n",
					key, err)
				failed = true
				break
			}
			item.data = nil
		}
		if failed {
			break
		} else {
			delete(ctx.deferredItems, key)
			iteration += 1
		}
	}
	if len(ctx.deferredItems) == 0 {
		stopTimer(ctx)
	}
	log.Infof("HandleDeferred() done map %d\n", len(ctx.deferredItems))
}

// Check if there are any deferred items for this key
func HasDeferred(key string) bool {
	if defaultCtx == nil {
		log.Fatal("HasDeferred no defaultCtx")
	}
	return defaultCtx.hasDeferred(key)
}

func (ctx *DeferredContext) hasDeferred(key string) bool {

	log.Debugf("HasDeferred(%s) map %d\n", key, len(ctx.deferredItems))
	_, ok := ctx.deferredItems[key]
	return ok
}

// Remove any item for the specific key. If no items left then stop timer.
func RemoveDeferred(key string) {
	if defaultCtx == nil {
		log.Fatal("RemoveDeferred no defaultCtx")
	}
	defaultCtx.removeDeferred(key)
}

func (ctx *DeferredContext) removeDeferred(key string) {

	log.Debugf("RemoveDeferred(%s) map %d\n", key, len(ctx.deferredItems))
	_, ok := ctx.deferredItems[key]
	if !ok {
		log.Errorf("removeDeferred: Non-existing key %s\n", key)
		return
	}
	log.Debugf("Deleting key %s\n", key)
	delete(ctx.deferredItems, key)

	if len(ctx.deferredItems) == 0 {
		stopTimer(ctx)
	}
}

// Replace any item for the specified key. If timer not running start it
func SetDeferred(key string, data []byte, url string,
	zedcloudCtx ZedCloudContext, ignore400 bool) {

	if defaultCtx == nil {
		log.Fatal("SetDeferred no defaultCtx")
	}
	defaultCtx.setDeferred(key, data, url, zedcloudCtx, ignore400)
}

func (ctx *DeferredContext) setDeferred(key string, data []byte, url string,
	zedcloudCtx ZedCloudContext, ignore400 bool) {

	log.Infof("SetDeferred(%s) map %d\n", key, len(ctx.deferredItems))
	if len(ctx.deferredItems) == 0 {
		startTimer(ctx)
	}
	_, ok := ctx.deferredItems[key]
	if ok {
		log.Debugf("Replacing key %s\n", key)
	} else {
		log.Debugf("Adding key %s\n", key)
	}
	item := deferredItem{
		data:        data,
		url:         url,
		zedcloudCtx: zedcloudCtx,
		ignore400:   ignore400,
	}
	l := deferredItemList{}
	l.list = append(l.list, item)
	ctx.deferredItems[key] = l
}

// Add to slice for this key
func AddDeferred(key string, data []byte, url string,
	zedcloudCtx ZedCloudContext, ignore400 bool) {

	if defaultCtx == nil {
		log.Fatal("SetDeferred no defaultCtx")
	}
	defaultCtx.addDeferred(key, data, url, zedcloudCtx, ignore400)
}

func (ctx *DeferredContext) addDeferred(key string, data []byte, url string,
	zedcloudCtx ZedCloudContext, ignore400 bool) {

	log.Infof("AddDeferred(%s) map %d\n", key, len(ctx.deferredItems))
	if len(ctx.deferredItems) == 0 {
		startTimer(ctx)
	}
	l, ok := ctx.deferredItems[key]
	if ok {
		log.Debugf("Appening to key %s len %d\n", key, len(l.list))
	} else {
		log.Debugf("Adding key %s\n", key)
	}
	item := deferredItem{
		data:        data,
		url:         url,
		zedcloudCtx: zedcloudCtx,
		ignore400:   ignore400,
	}
	l.list = append(l.list, item)
	ctx.deferredItems[key] = l
}

// Try every minute backoff to every 15 minutes
func startTimer(ctx *DeferredContext) {

	log.Debugf("startTimer()\n")
	min := 1 * time.Minute
	max := 15 * time.Minute
	ctx.ticker.UpdateExpTicker(min, max, 0.3)
}

func stopTimer(ctx *DeferredContext) {

	log.Debugf("stopTimer()\n")
	ctx.ticker.UpdateRangeTicker(longTime1, longTime2)
}
