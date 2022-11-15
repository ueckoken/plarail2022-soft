# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import block_pb2 as block__pb2


class BlockStateSyncStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.NotifyState = channel.unary_unary(
            "/BlockStateSync/NotifyState",
            request_serializer=block__pb2.NotifyStateRequest.SerializeToString,
            response_deserializer=block__pb2.NotifyStateResponse.FromString,
        )


class BlockStateSyncServicer(object):
    """Missing associated documentation comment in .proto file."""

    def NotifyState(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details("Method not implemented!")
        raise NotImplementedError("Method not implemented!")


def add_BlockStateSyncServicer_to_server(servicer, server):
    rpc_method_handlers = {
        "NotifyState": grpc.unary_unary_rpc_method_handler(
            servicer.NotifyState,
            request_deserializer=block__pb2.NotifyStateRequest.FromString,
            response_serializer=block__pb2.NotifyStateResponse.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        "BlockStateSync", rpc_method_handlers
    )
    server.add_generic_rpc_handlers((generic_handler,))


# This class is part of an EXPERIMENTAL API.
class BlockStateSync(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def NotifyState(
        request,
        target,
        options=(),
        channel_credentials=None,
        call_credentials=None,
        insecure=False,
        compression=None,
        wait_for_ready=None,
        timeout=None,
        metadata=None,
    ):
        return grpc.experimental.unary_unary(
            request,
            target,
            "/BlockStateSync/NotifyState",
            block__pb2.NotifyStateRequest.SerializeToString,
            block__pb2.NotifyStateResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
        )
