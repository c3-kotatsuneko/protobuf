// @generated by protoc-gen-es v2.0.0
// @generated from file game/game.proto (package game, syntax proto3)
/* eslint-disable */

import type { GenFile, GenService } from "@bufbuild/protobuf/codegenv1";
import type { GameStatusRequestSchema, GameStatusResponseSchema, PhysicsRequestSchema, PhysicsResponseSchema } from "./rpc/game_pb";

/**
 * Describes the file game/game.proto.
 */
export declare const file_game_game: GenFile;

/**
 * @generated from service game.GameService
 */
export declare const GameService: GenService<{
  /**
   * @generated from rpc game.GameService.GameStatusStream
   */
  gameStatusStream: {
    methodKind: "unary";
    input: typeof GameStatusRequestSchema;
    output: typeof GameStatusResponseSchema;
  },
  /**
   * @generated from rpc game.GameService.PhysicsStream
   */
  physicsStream: {
    methodKind: "unary";
    input: typeof PhysicsRequestSchema;
    output: typeof PhysicsResponseSchema;
  },
}>;

