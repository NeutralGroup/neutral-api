python_grpcio_library(
    name='user-defs-python',
    sources=['quotes.proto', 'instruments.proto', 'common.proto', 'accounts.proto'],
    dependencies=[
        '3rdparty/python:protobuf',
    ]
)

python_grpcio_library(
    name='user-api-python',
    sources=['neutralservices.proto'],
    dependencies=[
        ':user-defs-python',
        '3rdparty/python:protobuf',
    ]
)
