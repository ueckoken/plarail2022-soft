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


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0b\x62lock.proto\x1a\x1bgoogle/protobuf/empty.proto\"3\n\x0e\x42lockStateList\x12!\n\x0c\x62lock_states\x18\x01 \x03(\x0b\x32\x0b.BlockState\"\xe4\x04\n\nBlockState\x12$\n\x07\x62lockId\x18\x03 \x01(\x0e\x32\x13.BlockState.BlockId\x12 \n\x05state\x18\x02 \x01(\x0e\x32\x11.BlockState.State\"\xe2\x03\n\x07\x42lockId\x12\x0b\n\x07unknown\x10\x00\x12\x1b\n\x17shinjuku_sakurajosui_up\x10\n\x12\x1d\n\x19shinjuku_sakurajosui_down\x10\x0b\x12\x18\n\x14sakurajosui_chofu_up\x10\x14\x12\x1a\n\x16sakurajosui_chofu_down\x10\x15\x12\x15\n\x11\x63hofu_hachioji_up\x10\x1e\x12\x17\n\x13\x63hofu_hachioji_down\x10\x1f\x12\x16\n\x12\x63hofu_hashimoto_up\x10(\x12\x18\n\x14\x63hofu_hashimoto_down\x10)\x12\x0f\n\x0bshinjuku_b1\x10\x64\x12\x0f\n\x0bshinjuku_b2\x10\x65\x12\x12\n\x0esakurajosui_b1\x10n\x12\x12\n\x0esakurajosui_b2\x10o\x12\x12\n\x0esakurajosui_b3\x10x\x12\x12\n\x0esakurajosui_b4\x10y\x12\r\n\x08\x63hofu_b1\x10\x82\x01\x12\r\n\x08\x63hofu_b2\x10\x83\x01\x12\r\n\x08\x63hofu_b3\x10\x84\x01\x12\r\n\x08\x63hofu_b4\x10\x85\x01\x12\x11\n\x0chashimoto_b1\x10\x8c\x01\x12\x11\n\x0chashimoto_b2\x10\x8d\x01\x12\x10\n\x0bhachioji_b1\x10\x96\x01\x12\x10\n\x0bhachioji_b2\x10\x97\x01\")\n\x05State\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x08\n\x04OPEN\x10\x01\x12\t\n\x05\x43LOSE\x10\x02\x32J\n\x0e\x42lockStateSync\x12\x38\n\x0bNotifyState\x12\x16.google.protobuf.Empty\x1a\x0f.BlockStateList\"\x00\x42\x08Z\x06./specb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'block_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\006./spec'
  _BLOCKSTATELIST._serialized_start=44
  _BLOCKSTATELIST._serialized_end=95
  _BLOCKSTATE._serialized_start=98
  _BLOCKSTATE._serialized_end=710
  _BLOCKSTATE_BLOCKID._serialized_start=185
  _BLOCKSTATE_BLOCKID._serialized_end=667
  _BLOCKSTATE_STATE._serialized_start=669
  _BLOCKSTATE_STATE._serialized_end=710
  _BLOCKSTATESYNC._serialized_start=712
  _BLOCKSTATESYNC._serialized_end=786
# @@protoc_insertion_point(module_scope)
