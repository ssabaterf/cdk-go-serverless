import * as cdk from "aws-cdk-lib";
import { LambdaIntegration, RestApi } from "aws-cdk-lib/aws-apigateway";
import * as lambda from "aws-cdk-lib/aws-lambda";
import { Construct } from "constructs";

export class GoBEStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // The code that defines your stack goes here
    const regionInfoF = new lambda.Function(this, "region-info", {
      runtime: lambda.Runtime.GO_1_X,
      handler: "main",
      code: lambda.Code.fromAsset("src/region_info"),
    });

    const pingF = new lambda.Function(this, "ping", {
      runtime: lambda.Runtime.GO_1_X,
      handler: "main",
      code: lambda.Code.fromAsset("src/ping"),
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
