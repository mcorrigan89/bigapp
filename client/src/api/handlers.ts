import { ConnectError } from "@connectrpc/connect";
import { ErrorCode, ErrorDetails } from "@/api/gen/user/v1/user_pb";

export type ErrorHandler = {
  [key in ErrorCode]?: (message: string) => void;
};

type ServiceResult<T> = {
  data?: T;
  error?: ErrorDetails;
};

export async function handleServiceCall<T extends { error?: ErrorDetails }>(
  call: Promise<T>,
  errorHandlers: ErrorHandler,
): Promise<ServiceResult<T>> {
  try {
    const response = await call;
    if (response.error?.code && errorHandlers[response.error.code]) {
      return { error: response.error };
    }
    return { data: response };
  } catch (err) {
    if (err instanceof ConnectError) {
      return {
        error: {
          code: ErrorCode.UNSPECIFIED,
          message: err.message,
          $typeName: "user.v1.ErrorDetails",
        },
      };
    }
    throw err;
  }
}
