# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: config/netconfig.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from config import acipherinfo_pb2 as config_dot_acipherinfo__pb2
from config import fw_pb2 as config_dot_fw__pb2
from config import netcmn_pb2 as config_dot_netcmn__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x16\x63onfig/netconfig.proto\x12\x15org.lfedge.eve.config\x1a\x18\x63onfig/acipherinfo.proto\x1a\x0f\x63onfig/fw.proto\x1a\x13\x63onfig/netcmn.proto\"\x9f\x02\n\rNetworkConfig\x12\n\n\x02id\x18\x01 \x01(\t\x12\x30\n\x04type\x18\x05 \x01(\x0e\x32\".org.lfedge.eve.config.NetworkType\x12)\n\x02ip\x18\x06 \x01(\x0b\x32\x1d.org.lfedge.eve.config.ipspec\x12\x36\n\x03\x64ns\x18\x07 \x03(\x0b\x32).org.lfedge.eve.config.ZnetStaticDNSEntry\x12\x34\n\x08\x65ntProxy\x18\x08 \x01(\x0b\x32\".org.lfedge.eve.config.ProxyConfig\x12\x37\n\x08wireless\x18\n \x01(\x0b\x32%.org.lfedge.eve.config.WirelessConfig\"\xf9\x01\n\x0eNetworkAdapter\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x11\n\tnetworkId\x18\x03 \x01(\t\x12\x0c\n\x04\x61\x64\x64r\x18\x04 \x01(\t\x12\x10\n\x08hostname\x18\x05 \x01(\t\x12\x11\n\tcryptoEid\x18\n \x01(\t\x12\x15\n\rlispsignature\x18\x06 \x01(\t\x12\x0f\n\x07pemcert\x18\x07 \x01(\x0c\x12\x15\n\rpemprivatekey\x18\x08 \x01(\x0c\x12\x12\n\nmacAddress\x18\t \x01(\t\x12(\n\x04\x61\x63ls\x18( \x03(\x0b\x32\x1a.org.lfedge.eve.config.ACE\x12\x16\n\x0e\x61\x63\x63\x65ss_vlan_id\x18) \x01(\r\"\xb3\x01\n\x0eWirelessConfig\x12\x31\n\x04type\x18\x01 \x01(\x0e\x32#.org.lfedge.eve.config.WirelessType\x12:\n\x0b\x63\x65llularCfg\x18\x05 \x03(\x0b\x32%.org.lfedge.eve.config.CellularConfig\x12\x32\n\x07wifiCfg\x18\n \x03(\x0b\x32!.org.lfedge.eve.config.WifiConfig\"y\n\x0e\x43\x65llularConfig\x12\x0b\n\x03\x41PN\x18\x01 \x01(\t\x12?\n\x05probe\x18\x02 \x01(\x0b\x32\x30.org.lfedge.eve.config.CellularConnectivityProbe\x12\x19\n\x11location_tracking\x18\x03 \x01(\x08\"C\n\x19\x43\x65llularConnectivityProbe\x12\x0f\n\x07\x64isable\x18\x01 \x01(\x08\x12\x15\n\rprobe_address\x18\x02 \x01(\t\"\xb7\x02\n\nWifiConfig\x12\x10\n\x08wifiSSID\x18\x01 \x01(\t\x12\x37\n\tkeyScheme\x18\x02 \x01(\x0e\x32$.org.lfedge.eve.config.WiFiKeyScheme\x12\x10\n\x08identity\x18\x05 \x01(\t\x12\x10\n\x08password\x18\n \x01(\t\x12=\n\x06\x63rypto\x18\x14 \x01(\x0b\x32-.org.lfedge.eve.config.WifiConfig.cryptoblock\x12\x10\n\x08priority\x18\x19 \x01(\x05\x12\x36\n\ncipherData\x18\x1e \x01(\x0b\x32\".org.lfedge.eve.config.CipherBlock\x1a\x31\n\x0b\x63ryptoblock\x12\x10\n\x08identity\x18\x0b \x01(\t\x12\x10\n\x08password\x18\x0c \x01(\tB=\n\x15org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/configb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'config.netconfig_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\025org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/config'
  _NETWORKCONFIG._serialized_start=114
  _NETWORKCONFIG._serialized_end=401
  _NETWORKADAPTER._serialized_start=404
  _NETWORKADAPTER._serialized_end=653
  _WIRELESSCONFIG._serialized_start=656
  _WIRELESSCONFIG._serialized_end=835
  _CELLULARCONFIG._serialized_start=837
  _CELLULARCONFIG._serialized_end=958
  _CELLULARCONNECTIVITYPROBE._serialized_start=960
  _CELLULARCONNECTIVITYPROBE._serialized_end=1027
  _WIFICONFIG._serialized_start=1030
  _WIFICONFIG._serialized_end=1341
  _WIFICONFIG_CRYPTOBLOCK._serialized_start=1292
  _WIFICONFIG_CRYPTOBLOCK._serialized_end=1341
# @@protoc_insertion_point(module_scope)
