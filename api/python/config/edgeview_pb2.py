# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: config/edgeview.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='config/edgeview.proto',
  package='org.lfedge.eve.config',
  syntax='proto3',
  serialized_options=b'\n\025org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/config',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x15\x63onfig/edgeview.proto\x12\x15org.lfedge.eve.config\"\xb8\x01\n\x0e\x45\x64geViewConfig\x12\r\n\x05token\x18\x01 \x01(\t\x12\x15\n\rdisp_cert_pem\x18\x02 \x03(\x0c\x12?\n\ndev_policy\x18\x03 \x01(\x0b\x32+.org.lfedge.eve.config.DevDebugAccessPolicy\x12?\n\napp_policy\x18\x04 \x01(\x0b\x32+.org.lfedge.eve.config.AppDebugAccessPolicy\")\n\x14\x44\x65vDebugAccessPolicy\x12\x11\n\tallow_dev\x18\x01 \x01(\x08\")\n\x14\x41ppDebugAccessPolicy\x12\x11\n\tallow_app\x18\x01 \x01(\x08\x42=\n\x15org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/configb\x06proto3'
)




_EDGEVIEWCONFIG = _descriptor.Descriptor(
  name='EdgeViewConfig',
  full_name='org.lfedge.eve.config.EdgeViewConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='token', full_name='org.lfedge.eve.config.EdgeViewConfig.token', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='disp_cert_pem', full_name='org.lfedge.eve.config.EdgeViewConfig.disp_cert_pem', index=1,
      number=2, type=12, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='dev_policy', full_name='org.lfedge.eve.config.EdgeViewConfig.dev_policy', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='app_policy', full_name='org.lfedge.eve.config.EdgeViewConfig.app_policy', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=49,
  serialized_end=233,
)


_DEVDEBUGACCESSPOLICY = _descriptor.Descriptor(
  name='DevDebugAccessPolicy',
  full_name='org.lfedge.eve.config.DevDebugAccessPolicy',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='allow_dev', full_name='org.lfedge.eve.config.DevDebugAccessPolicy.allow_dev', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=235,
  serialized_end=276,
)


_APPDEBUGACCESSPOLICY = _descriptor.Descriptor(
  name='AppDebugAccessPolicy',
  full_name='org.lfedge.eve.config.AppDebugAccessPolicy',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='allow_app', full_name='org.lfedge.eve.config.AppDebugAccessPolicy.allow_app', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=278,
  serialized_end=319,
)

_EDGEVIEWCONFIG.fields_by_name['dev_policy'].message_type = _DEVDEBUGACCESSPOLICY
_EDGEVIEWCONFIG.fields_by_name['app_policy'].message_type = _APPDEBUGACCESSPOLICY
DESCRIPTOR.message_types_by_name['EdgeViewConfig'] = _EDGEVIEWCONFIG
DESCRIPTOR.message_types_by_name['DevDebugAccessPolicy'] = _DEVDEBUGACCESSPOLICY
DESCRIPTOR.message_types_by_name['AppDebugAccessPolicy'] = _APPDEBUGACCESSPOLICY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

EdgeViewConfig = _reflection.GeneratedProtocolMessageType('EdgeViewConfig', (_message.Message,), {
  'DESCRIPTOR' : _EDGEVIEWCONFIG,
  '__module__' : 'config.edgeview_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.config.EdgeViewConfig)
  })
_sym_db.RegisterMessage(EdgeViewConfig)

DevDebugAccessPolicy = _reflection.GeneratedProtocolMessageType('DevDebugAccessPolicy', (_message.Message,), {
  'DESCRIPTOR' : _DEVDEBUGACCESSPOLICY,
  '__module__' : 'config.edgeview_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.config.DevDebugAccessPolicy)
  })
_sym_db.RegisterMessage(DevDebugAccessPolicy)

AppDebugAccessPolicy = _reflection.GeneratedProtocolMessageType('AppDebugAccessPolicy', (_message.Message,), {
  'DESCRIPTOR' : _APPDEBUGACCESSPOLICY,
  '__module__' : 'config.edgeview_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.config.AppDebugAccessPolicy)
  })
_sym_db.RegisterMessage(AppDebugAccessPolicy)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
