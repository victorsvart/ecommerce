import { z } from "zod";

export class ValidationError extends Error {
  constructor(message: string) {
    super(message);
    this.cause = 400;
    this.name = "ValidationError";
  }
}

/**
 * -- Victor
 * A global singleton function for advanced validation of form values.
 *
 * This function takes a Zod schema (`type`) and a `FormData` object, validates
 * the form data against the schema, and throws a `ValidationError` if the validation fails.
 * It returns the parsed data if the validation is successful.
 *
 * @param {z.ZodType<T, any, any>} type - A Zod schema that defines the expected shape of the form data.
 *   The schema is used to validate the `formData`.
 *
 * @param {FormData} formData - The `FormData` object containing the raw form data that needs to be validated.
 *
 * @returns {Promise<z.SafeParseSuccess<T>>} - A promise that resolves to a `SafeParseSuccess` object containing the validated form data.
 *   If the validation succeeds, `result.data` will contain the parsed and validated form data.
 *
 * @throws {ValidationError} - Throws a custom `ValidationError` if the validation fails, with the error message being
 *   extracted from the validation errors returned by Zod.
 *
 * @template T - The type that the form data will be validated and parsed into.
 */ export async function ValidateForm<T>(
  type: z.ZodType<T, any, any>,
  formData: FormData
): Promise<z.SafeParseSuccess<T>> {
  const formValues = Object.fromEntries(formData);
  const result = await type.safeParseAsync(formValues);
  if (!result.success) {
    const errors = result.error.flatten();
    const errorMessage =
      Object.values(errors.fieldErrors).flat()[0] || "Invalid form data";
    throw new ValidationError(errorMessage);
  }

  return result;
}
