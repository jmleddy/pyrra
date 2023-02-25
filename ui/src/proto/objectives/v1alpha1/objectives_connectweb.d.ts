// @generated by protoc-gen-connect-web v0.8.1 with parameter "target=js+dts"
// @generated from file objectives/v1alpha1/objectives.proto (package objectives.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetAlertsRequest, GetAlertsResponse, GetStatusRequest, GetStatusResponse, GraphDurationRequest, GraphDurationResponse, GraphErrorBudgetRequest, GraphErrorBudgetResponse, GraphErrorsRequest, GraphErrorsResponse, GraphRateRequest, GraphRateResponse, ListRequest, ListResponse } from "./objectives_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service objectives.v1alpha1.ObjectiveService
 */
export declare const ObjectiveService: {
  readonly typeName: "objectives.v1alpha1.ObjectiveService",
  readonly methods: {
    /**
     * @generated from rpc objectives.v1alpha1.ObjectiveService.List
     */
    readonly list: {
      readonly name: "List",
      readonly I: typeof ListRequest,
      readonly O: typeof ListResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc objectives.v1alpha1.ObjectiveService.GetStatus
     */
    readonly getStatus: {
      readonly name: "GetStatus",
      readonly I: typeof GetStatusRequest,
      readonly O: typeof GetStatusResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc objectives.v1alpha1.ObjectiveService.GetAlerts
     */
    readonly getAlerts: {
      readonly name: "GetAlerts",
      readonly I: typeof GetAlertsRequest,
      readonly O: typeof GetAlertsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc objectives.v1alpha1.ObjectiveService.GraphErrorBudget
     */
    readonly graphErrorBudget: {
      readonly name: "GraphErrorBudget",
      readonly I: typeof GraphErrorBudgetRequest,
      readonly O: typeof GraphErrorBudgetResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc objectives.v1alpha1.ObjectiveService.GraphRate
     */
    readonly graphRate: {
      readonly name: "GraphRate",
      readonly I: typeof GraphRateRequest,
      readonly O: typeof GraphRateResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc objectives.v1alpha1.ObjectiveService.GraphErrors
     */
    readonly graphErrors: {
      readonly name: "GraphErrors",
      readonly I: typeof GraphErrorsRequest,
      readonly O: typeof GraphErrorsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc objectives.v1alpha1.ObjectiveService.GraphDuration
     */
    readonly graphDuration: {
      readonly name: "GraphDuration",
      readonly I: typeof GraphDurationRequest,
      readonly O: typeof GraphDurationResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

/**
 * @generated from service objectives.v1alpha1.ObjectiveBackendService
 */
export declare const ObjectiveBackendService: {
  readonly typeName: "objectives.v1alpha1.ObjectiveBackendService",
  readonly methods: {
    /**
     * @generated from rpc objectives.v1alpha1.ObjectiveBackendService.List
     */
    readonly list: {
      readonly name: "List",
      readonly I: typeof ListRequest,
      readonly O: typeof ListResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

