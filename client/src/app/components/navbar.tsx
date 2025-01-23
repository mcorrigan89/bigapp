"use client";

import Link from "next/link";
import {
  Dialog,
  DialogTrigger,
  Heading,
  Input,
  Label,
  Modal,
  ModalOverlay,
  TextField,
  Disclosure,
  DisclosurePanel,
} from "react-aria-components";
import { Button } from "@/components/button";
import { cn } from "@/lib/utils";

export function Navbar() {
  return (
    <nav className="bg-white shadow-sm">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="flex h-16 justify-between">
          <div className="flex">
            <div className="flex shrink-0 items-center">
              <img
                alt="Your Company"
                src="https://tailwindui.com/plus/img/logos/mark.svg?color=indigo&shade=600"
                className="h-8 w-auto"
              />
            </div>
            <div className="hidden sm:ml-6 sm:flex sm:space-x-8">
              {/* Current: "border-indigo-500 text-gray-900", Default: "border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700" */}
              <a
                href="#"
                className="border-indigo-500 text-gray-900 inline-flex items-center border-b-2 px-1 pt-1 text-sm font-medium"
              >
                Dashboard
              </a>
              <a
                href="#"
                className="text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center border-b-2 border-transparent px-1 pt-1 text-sm font-medium"
              >
                Team
              </a>
              <a
                href="#"
                className="text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center border-b-2 border-transparent px-1 pt-1 text-sm font-medium"
              >
                Projects
              </a>
              <a
                href="#"
                className="text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center border-b-2 border-transparent px-1 pt-1 text-sm font-medium"
              >
                Calendar
              </a>
            </div>
          </div>
          <div className="hidden sm:ml-6 sm:flex sm:items-center">
            {/* Profile dropdown */}
            <div className="relative ml-3">
              <div>
                <Button className="focus:ring-indigo-500 relative flex rounded-full bg-white text-sm focus:ring-2 focus:ring-offset-2 focus:outline-hidden">
                  <span className="absolute -inset-1.5" />
                  <span className="sr-only">Open user menu</span>
                  <img
                    alt=""
                    src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
                    className="size-8 rounded-full"
                  />
                </Button>
              </div>
              <ul className="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 ring-1 shadow-lg ring-black/5 transition focus:outline-hidden data-closed:scale-95 data-closed:transform data-closed:opacity-0 data-enter:duration-200 data-enter:ease-out data-leave:duration-75 data-leave:ease-in">
                <li>
                  <a
                    href="#"
                    className="text-gray-700 data-focus:bg-gray-100 block px-4 py-2 text-sm data-focus:outline-hidden"
                  >
                    Your Profile
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="text-gray-700 data-focus:bg-gray-100 block px-4 py-2 text-sm data-focus:outline-hidden"
                  >
                    Settings
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="text-gray-700 data-focus:bg-gray-100 block px-4 py-2 text-sm data-focus:outline-hidden"
                  >
                    Sign out
                  </a>
                </li>
              </ul>
            </div>
          </div>
          <div className="-mr-2 flex items-center sm:hidden">
            {/* Mobile menu button */}
            <Button className="group text-gray-400 hover:bg-gray-100 hover:text-gray-500 focus:ring-indigo-500 relative inline-flex items-center justify-center rounded-md p-2 focus:ring-2 focus:outline-hidden focus:ring-inset">
              <span className="absolute -inset-0.5" />
              <span className="sr-only">Open main menu</span>
            </Button>
          </div>
        </div>
      </div>
    </nav>
  );
}
