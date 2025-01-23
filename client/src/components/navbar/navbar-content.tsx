"use client";
import React, { useState } from "react";
import {
  Menu,
  MenuItem,
  MenuTrigger,
  Button,
  Popover,
} from "react-aria-components";
import { ChevronDown, Menu as MenuIcon, X } from "lucide-react";
import { User } from "@/api/gen/user/v1/user_pb";
import { Avatar } from "../avatar";

interface NavBarContentProps {
  links: { href: string; label: string }[];
  currentUser: User | null;
}
export const NavBarContent = ({ links, currentUser }: NavBarContentProps) => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  return (
    <nav className="bg-white shadow">
      <div className="mx-auto max-w-7xl px-4">
        <div className="flex h-16 items-center justify-between">
          {/* Logo */}
          <div className="flex-shrink-0">
            <span className="text-xl font-bold">Logo</span>
          </div>

          {/* Mobile Menu Button */}
          <div className="md:hidden">
            <button
              onClick={() => setIsMenuOpen(!isMenuOpen)}
              className="inline-flex items-center justify-center rounded-md p-2 text-gray-400 hover:bg-gray-100 hover:text-gray-500 focus:outline-none"
              aria-expanded={isMenuOpen}
            >
              <span className="sr-only">Open main menu</span>
              {isMenuOpen ? (
                <X className="block h-6 w-6" aria-hidden="true" />
              ) : (
                <MenuIcon className="block h-6 w-6" aria-hidden="true" />
              )}
            </button>
          </div>

          {/* Desktop Navigation */}
          <div className="hidden space-x-8 md:flex">
            {links.map((link) => (
              <a
                key={link.href}
                href={link.href}
                className="px-3 py-2 text-sm font-medium text-gray-500 hover:text-gray-900"
              >
                {link.label}
              </a>
            ))}
          </div>

          {/* Profile Menu */}
          <div className="hidden md:block">
            <MenuTrigger>
              <Button
                className="flex items-center space-x-2 focus:outline-none"
                aria-label="User menu"
              >
                {/* <div className="h-8 w-8 overflow-hidden rounded-full">
                  <img
                    src="/api/placeholder/32/32"
                    alt="User avatar"
                    className="h-full w-full object-cover"
                  />
                </div> */}
                <Avatar avatar={currentUser?.avatar} size={"small"} />
                <ChevronDown className="h-4 w-4 text-gray-500" />
              </Button>
              <Popover className="mt-2">
                <Menu className="min-w-48 rounded-md border bg-white py-1 shadow-lg">
                  <MenuItem className="cursor-pointer px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">
                    Your Profile
                  </MenuItem>
                  <MenuItem className="cursor-pointer px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">
                    Settings
                  </MenuItem>
                  <MenuItem className="cursor-pointer px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">
                    Sign out
                  </MenuItem>
                </Menu>
              </Popover>
            </MenuTrigger>
          </div>
        </div>
      </div>

      {/* Mobile Menu */}
      <div className={`${isMenuOpen ? "block" : "hidden"} md:hidden`}>
        <div className="space-y-1 px-2 pt-2 pb-3">
          {links.map((link) => (
            <a
              key={link.href}
              href={link.href}
              className="block px-3 py-2 text-base font-medium text-gray-700 hover:bg-gray-50 hover:text-gray-900"
            >
              {link.label}
            </a>
          ))}
        </div>
        {/* Mobile Profile Menu */}
        <div className="border-t border-gray-200 pt-4 pb-3">
          <div className="flex items-center px-4">
            <div className="flex-shrink-0">
              <Avatar avatar={currentUser?.avatar} />
            </div>
            <div className="ml-3">
              <div className="text-base font-medium text-gray-800">
                User Name
              </div>
              <div className="text-sm font-medium text-gray-500">
                user@example.com
              </div>
            </div>
          </div>
          <div className="mt-3 space-y-1 px-2">
            <a
              href="#"
              className="block px-3 py-2 text-base font-medium text-gray-700 hover:bg-gray-50 hover:text-gray-900"
            >
              Your Profile
            </a>
            <a
              href="#"
              className="block px-3 py-2 text-base font-medium text-gray-700 hover:bg-gray-50 hover:text-gray-900"
            >
              Settings
            </a>
            <a
              href="#"
              className="block px-3 py-2 text-base font-medium text-gray-700 hover:bg-gray-50 hover:text-gray-900"
            >
              Sign out
            </a>
          </div>
        </div>
      </div>
    </nav>
  );
};
