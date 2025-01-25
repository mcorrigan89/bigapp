"use client";

import { useState } from "react";
import { Dialog, DialogBackdrop, DialogPanel, DialogTitle } from "@headlessui/react";
import { CheckIcon } from "@heroicons/react/24/outline";
import { useHookFormAction } from "@next-safe-action/adapter-react-hook-form/hooks";
import { createCollectionAction } from "./action";
import { createCollectionSchema } from "./schema";
import { zodResolver } from "@hookform/resolvers/zod";

export default function CreateCollectionModal() {
  const [open, setOpen] = useState(false);

  const { form, handleSubmitWithAction } = useHookFormAction(
    createCollectionAction,
    zodResolver(createCollectionSchema),
    {
      actionProps: {
        onSuccess: () => {
          setOpen(false);
        },
      },
      formProps: {
        defaultValues: {},
      },
      errorMapProps: {},
    },
  );

  const errors = form.formState.errors;

  return (
    <>
      <button
        onClick={() => setOpen(!open)}
        type="button"
        className="ml-3 inline-flex items-center rounded-md bg-indigo-500 px-3 py-2 text-sm font-semibold text-white shadow-xs hover:bg-indigo-400 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500"
      >
        Create Collection
      </button>
      <Dialog open={open} onClose={setOpen} className="relative z-50">
        <DialogBackdrop
          transition
          className="fixed inset-0 bg-gray-950/75 transition-opacity data-closed:opacity-0 data-enter:duration-300 data-enter:ease-out data-leave:duration-200 data-leave:ease-in"
        />

        <div className="fixed inset-0 z-10 w-screen overflow-y-auto">
          <div className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
            <DialogPanel
              transition
              className="relative transform overflow-hidden rounded-lg bg-gray-800 px-4 pt-5 pb-4 text-left shadow-xl transition-all data-closed:translate-y-4 data-closed:opacity-0 data-enter:duration-300 data-enter:ease-out data-leave:duration-200 data-leave:ease-in sm:my-8 sm:w-full sm:max-w-sm sm:p-6 data-closed:sm:translate-y-0 data-closed:sm:scale-95"
            >
              <form onSubmit={handleSubmitWithAction}>
                <div>
                  <div className="mx-auto flex size-12 items-center justify-center rounded-full bg-green-600">
                    <CheckIcon aria-hidden="true" className="size-6 text-green-100" />
                  </div>
                  <div className="mt-3 text-center sm:mt-5">
                    <DialogTitle as="h3" className="text-base font-semibold text-white">
                      Create a new photo collection
                    </DialogTitle>
                    <div className="mt-2">
                      <div className="sm:col-span-3">
                        <label htmlFor="first-name" className="block text-sm/6 font-medium text-white">
                          First name
                        </label>
                        <div className="mt-2">
                          <input
                            {...form.register("collectionName")}
                            type="text"
                            autoComplete="given-name"
                            className="block w-full rounded-md bg-white/5 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 placeholder:text-gray-500 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-500 sm:text-sm/6"
                          />
                        </div>
                        {errors.collectionName?.message ? (
                          <p id="givenName-error" className="mt-2 text-sm text-red-600">
                            {errors.collectionName?.message}
                          </p>
                        ) : null}
                      </div>
                    </div>
                  </div>
                </div>
                <div className="mt-5 sm:mt-6">
                  <button
                    type="submit"
                    className="inline-flex w-full justify-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                  >
                    Go back to dashboard
                  </button>
                </div>
              </form>
            </DialogPanel>
          </div>
        </div>
      </Dialog>
    </>
  );
}
