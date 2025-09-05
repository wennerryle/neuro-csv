import * as v from "npm:valibot";
import { toJsonSchema } from "npm:@valibot/to-json-schema";

const ArgumentsSchema = v.object({
  model: v.pipe(
    v.string(),
    v.nonEmpty(),
    v.description("LM Studio model identifier"),
  ),
  "base-url": v.optional(
    v.pipe(
      v.string(),
      v.nonEmpty(),
      v.url(),
      v.description("URL of LM Studio server"),
    ),
    "ws://localhost:1234",
  ),
  "structured-output-file": v.pipe(
    v.string(),
    v.nonEmpty(),
    v.description(
      "A path to file that contains JSON Schema for structured output",
    ),
  ),
  "instructions-file": v.pipe(
    v.string(),
    v.nonEmpty(),
    v.description("A path to file that contains instructions for AI"),
  ),
  "slice": v.optional(v.pipe(
    v.array(v.union([
      v.pipe(
        v.string(),
        v.nonEmpty(),
        v.regex(/\d+\-\d+/),
      ),
      v.number(),
    ])),
    v.description(
      "Takes only rows specified in slice. Maybe be useful for debugging",
    ),
  )),
});

Deno.writeFileSync(
  "../neuro-csv.schema.json",
  new TextEncoder().encode(
    JSON.stringify(toJsonSchema(ArgumentsSchema), null, 2),
  ),
);
