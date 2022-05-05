var PROTO_PATH = __dirname + '/../../greeter.proto';
var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
const { assert, expect, should } = require('chai')

var client


beforeEach(function () {
    var packageDefinition = protoLoader.loadSync(
        PROTO_PATH,
        {
            keepCase: true,
            longs: String,
            enums: String,
            defaults: true,
            oneofs: true,
            includeDirs: [
                // '/../third_party',
                __dirname + '/../node_modules/protobufjs'
            ]
        });
    var greeter_proto = grpc.loadPackageDefinition(packageDefinition).helloworld.v1;
    client = new greeter_proto.Greeter('localhost:9000', grpc.credentials.createInsecure());
})


describe("test_greeter", function () {
    it("test", function () {
        client.SayHello({ name: "test" }, function (err, resp) {
            assert.equal(err, null, "error should be null")
            assert.deepStrictEqual(resp, { message: "Hello test" })
        })
    })
})