# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: config/fw.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='config/fw.proto',
  package='org.lfedge.eve.config',
  syntax='proto3',
  serialized_options=b'\n\025org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/config',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0f\x63onfig/fw.proto\x12\x15org.lfedge.eve.config\"\'\n\x08\x41\x43\x45Match\x12\x0c\n\x04type\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t\"\x84\x01\n\tACEAction\x12\x0c\n\x04\x64rop\x18\x01 \x01(\x08\x12\r\n\x05limit\x18\x02 \x01(\x08\x12\x11\n\tlimitrate\x18\x03 \x01(\r\x12\x11\n\tlimitunit\x18\x04 \x01(\t\x12\x12\n\nlimitburst\x18\x05 \x01(\r\x12\x0f\n\x07portmap\x18\x06 \x01(\x08\x12\x0f\n\x07\x61ppPort\x18\x07 \x01(\r\"\xb6\x01\n\x03\x41\x43\x45\x12\x30\n\x07matches\x18\x01 \x03(\x0b\x32\x1f.org.lfedge.eve.config.ACEMatch\x12\x31\n\x07\x61\x63tions\x18\x02 \x03(\x0b\x32 .org.lfedge.eve.config.ACEAction\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\n\n\x02id\x18\x04 \x01(\x05\x12\x30\n\x03\x64ir\x18\x05 \x01(\x0e\x32#.org.lfedge.eve.config.ACEDirection*1\n\x0c\x41\x43\x45\x44irection\x12\x08\n\x04\x42OTH\x10\x00\x12\x0b\n\x07INGRESS\x10\x01\x12\n\n\x06\x45GRESS\x10\x02\x42=\n\x15org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/configb\x06proto3'
)

_ACEDIRECTION = _descriptor.EnumDescriptor(
  name='ACEDirection',
  full_name='org.lfedge.eve.config.ACEDirection',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='BOTH', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='INGRESS', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EGRESS', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=403,
  serialized_end=452,
)
_sym_db.RegisterEnumDescriptor(_ACEDIRECTION)

ACEDirection = enum_type_wrapper.EnumTypeWrapper(_ACEDIRECTION)
BOTH = 0
INGRESS = 1
EGRESS = 2



_ACEMATCH = _descriptor.Descriptor(
  name='ACEMatch',
  full_name='org.lfedge.eve.config.ACEMatch',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='org.lfedge.eve.config.ACEMatch.type', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='org.lfedge.eve.config.ACEMatch.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=42,
  serialized_end=81,
)


_ACEACTION = _descriptor.Descriptor(
  name='ACEAction',
  full_name='org.lfedge.eve.config.ACEAction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='drop', full_name='org.lfedge.eve.config.ACEAction.drop', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='limit', full_name='org.lfedge.eve.config.ACEAction.limit', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='limitrate', full_name='org.lfedge.eve.config.ACEAction.limitrate', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='limitunit', full_name='org.lfedge.eve.config.ACEAction.limitunit', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='limitburst', full_name='org.lfedge.eve.config.ACEAction.limitburst', index=4,
      number=5, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='portmap', full_name='org.lfedge.eve.config.ACEAction.portmap', index=5,
      number=6, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='appPort', full_name='org.lfedge.eve.config.ACEAction.appPort', index=6,
      number=7, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=84,
  serialized_end=216,
)


_ACE = _descriptor.Descriptor(
  name='ACE',
  full_name='org.lfedge.eve.config.ACE',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='matches', full_name='org.lfedge.eve.config.ACE.matches', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='actions', full_name='org.lfedge.eve.config.ACE.actions', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='org.lfedge.eve.config.ACE.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='id', full_name='org.lfedge.eve.config.ACE.id', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='dir', full_name='org.lfedge.eve.config.ACE.dir', index=4,
      number=5, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=219,
  serialized_end=401,
)

_ACE.fields_by_name['matches'].message_type = _ACEMATCH
_ACE.fields_by_name['actions'].message_type = _ACEACTION
_ACE.fields_by_name['dir'].enum_type = _ACEDIRECTION
DESCRIPTOR.message_types_by_name['ACEMatch'] = _ACEMATCH
DESCRIPTOR.message_types_by_name['ACEAction'] = _ACEACTION
DESCRIPTOR.message_types_by_name['ACE'] = _ACE
DESCRIPTOR.enum_types_by_name['ACEDirection'] = _ACEDIRECTION
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ACEMatch = _reflection.GeneratedProtocolMessageType('ACEMatch', (_message.Message,), {
  'DESCRIPTOR' : _ACEMATCH,
  '__module__' : 'config.fw_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.config.ACEMatch)
  })
_sym_db.RegisterMessage(ACEMatch)

ACEAction = _reflection.GeneratedProtocolMessageType('ACEAction', (_message.Message,), {
  'DESCRIPTOR' : _ACEACTION,
  '__module__' : 'config.fw_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.config.ACEAction)
  })
_sym_db.RegisterMessage(ACEAction)

ACE = _reflection.GeneratedProtocolMessageType('ACE', (_message.Message,), {
  'DESCRIPTOR' : _ACE,
  '__module__' : 'config.fw_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.config.ACE)
  })
_sym_db.RegisterMessage(ACE)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
