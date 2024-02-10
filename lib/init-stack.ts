import { GoFunction } from "@aws-cdk/aws-lambda-go-alpha";
import * as cdk from "aws-cdk-lib";
import { LambdaIntegration, RestApi } from "aws-cdk-lib/aws-apigateway";
import { Construct } from "constructs";

export class GoBEStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // The code that defines your stack goes here
    const regionInfoF = new GoFunction(this, "region-info", {
      entry: "src/region_info",
      bundling: {
        goBuildFlags: ['-ldflags "-s -w"'],
        environment: {
          GOOS: "linux",
          GOARCH: "amd64",
        },
      },
    });

    const pingF = new GoFunction(this, "ping", {
      entry: "src/ping",
      bundling: {
        goBuildFlags: ['-ldflags "-s -w"'],
        environment: {
          GOOS: "linux",
          GOARCH: "amd64",
        },
      },
    });

    const gateway = new RestApi(this, "rest-http-api", {
      defaultCorsPreflightOptions: {
        allowOrigins: ["*"],
        allowMethods: ["GET", "POST", "PUT", "DELETE", "OPTIONS"],
        allowHeaders: ["*"],
        allowCredentials: true,
      },
    });

    const refionInfoI = new LambdaIntegration(regionInfoF);
    const pingI = new LambdaIntegration(pingF);

    const regionInfoR = gateway.root.addResource("region-info");
    const pingR = gateway.root.addResource("ping");

    regionInfoR.addMethod("GET", refionInfoI);
    regionInfoR.addMethod("POST", refionInfoI);

    pingR.addMethod("GET", pingI);
  }
}
