"use client";

import qs from "query-string";
import { Button } from "./ui/button";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { useCallback } from "react";
import queryString from "query-string";

type Props = {
  label: any;
  selected?: boolean;
};

export const WordListCategorie = ({ selected, label }: Props) => {
  const router = useRouter();
  const searchParams = useSearchParams();

  const onClick = useCallback(() => {
    let currentQuery = {};

    if (searchParams) {
      currentQuery = queryString.parse(searchParams.toString());
    }

    const updatedQuery: any = {
      ...currentQuery,
      category: label,
    };

    if (searchParams?.get("category") === label) {
      delete updatedQuery.category;
    }

    const url = queryString.stringifyUrl(
      {
        url: "/explore",
        query: updatedQuery,
      },
      { skipNull: true }
    );

    router.push(url);
  }, [label, searchParams, router]);

  return (
    <Button
      onClick={onClick}
      type='button'
      className={` 'text-gray-900 border  border-white hover:border-gray-200 dark:border-gray-900 dark:bg-gray-900 dark:hover:border-gray-700 bg-white focus:ring-4 focus:outline-none focus:ring-gray-300 rounded-full text-base font-medium px-5 py-2.5 text-center me-3 mb-3 dark:text-white dark:focus:ring-gray-800' 
        
        ${selected ? "border-b-neutral-800" : "border-transparent"}
        ${selected ? "text-neutral-800" : "text-neutral-500"}`}
    >
      {label}
    </Button>
  );
};
