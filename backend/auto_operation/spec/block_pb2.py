# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: block.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0b\x62lock.proto\"\x8b\x01\n\x17NotifyBlockStateRequest\x12-\n\x05state\x18\x02 \x01(\x0e\x32\x1e.NotifyBlockStateRequest.State\x12\x16\n\x05\x62lock\x18\x03 \x01(\x0b\x32\x07.Blocks\")\n\x05State\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x08\n\x04OPEN\x10\x01\x12\t\n\x05\x43LOSE\x10\x02\"\x82\x01\n\x18NotifyBlockStateResponse\x12\x34\n\x08response\x18\x01 \x01(\x0e\x32\".NotifyBlockStateResponse.Response\"0\n\x08Response\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0b\n\x07SUCCESS\x10\x01\x12\n\n\x06\x46\x41ILED\x10\x02\"\xe9\x02\n\x06\x42locks\x12 \n\x07\x62lockId\x18\x03 \x01(\x0e\x32\x0f.Blocks.BlockId\"\xbc\x02\n\x07\x42lockId\x12\x0b\n\x07unknown\x10\x00\x12\x0f\n\x0bshinjuku_b1\x10\x01\x12\x0f\n\x0bshinjuku_b2\x10\x02\x12\x12\n\x0esakurajosui_b1\x10\x0b\x12\x12\n\x0esakurajosui_b2\x10\x0c\x12\x12\n\x0esakurajosui_b3\x10\r\x12\x12\n\x0esakurajosui_b4\x10\x0e\x12\x12\n\x0esakurajosui_b5\x10\x0f\x12\x12\n\x0esakurajosui_b6\x10\x10\x12\x0c\n\x08\x63hofu_b1\x10\x15\x12\x0c\n\x08\x63hofu_b2\x10\x16\x12\x0c\n\x08\x63hofu_b3\x10\x17\x12\x0c\n\x08\x63hofu_b4\x10\x18\x12\x0c\n\x08\x63hofu_b5\x10\x19\x12\x10\n\x0chashimoto_b1\x10\x1f\x12\x10\n\x0chashimoto_b2\x10 \x12\x0f\n\x0bhachioji_b1\x10)\x12\x0f\n\x0bhashioji_b2\x10*2V\n\x0e\x42lockStateSync\x12\x44\n\x0bNotifyState\x12\x18.NotifyBlockStateRequest\x1a\x19.NotifyBlockStateResponse\"\x00\x42\x08Z\x06./specb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'block_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\006./spec'
  _NOTIFYBLOCKSTATEREQUEST._serialized_start=16
  _NOTIFYBLOCKSTATEREQUEST._serialized_end=155
  _NOTIFYBLOCKSTATEREQUEST_STATE._serialized_start=114
  _NOTIFYBLOCKSTATEREQUEST_STATE._serialized_end=155
  _NOTIFYBLOCKSTATERESPONSE._serialized_start=158
  _NOTIFYBLOCKSTATERESPONSE._serialized_end=288
  _NOTIFYBLOCKSTATERESPONSE_RESPONSE._serialized_start=240
  _NOTIFYBLOCKSTATERESPONSE_RESPONSE._serialized_end=288
  _BLOCKS._serialized_start=291
  _BLOCKS._serialized_end=652
  _BLOCKS_BLOCKID._serialized_start=336
  _BLOCKS_BLOCKID._serialized_end=652
  _BLOCKSTATESYNC._serialized_start=654
  _BLOCKSTATESYNC._serialized_end=740
# @@protoc_insertion_point(module_scope)
