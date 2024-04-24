"use client";

import Image from "next/image";
import { useRouter, useSearchParams } from "next/navigation";
import queryString from "query-string";
import { useCallback } from "react";

type Props = {
  categorieAttribute: any;
  title: string;
};

export const CategorieCard = ({ title, categorieAttribute }: Props) => {
  const router = useRouter();
  const searchParams = useSearchParams();

  const onClick = useCallback(() => {
    let currentQuery = {};

    if (searchParams) {
      currentQuery = queryString.parse(searchParams.toString());
    }

    const updatedQuery: any = {
      ...currentQuery,
      category: title,
    };

    if (searchParams?.get("category") === title) {
      delete updatedQuery.category;
    }

    const url = queryString.stringifyUrl(
      {
        url: "explore/words",
        query: updatedQuery,
      },
      { skipNull: true }
    );

    router.push(url);
  }, [title, searchParams, router]);

  return (
    <div onClick={onClick} className='col-span-1 cursor-pointer group'>
      <div className=' flex flex-col gap-2 w-full'>
        <div className='overflow-hidden w-full rounded-xl '>
          <Image
            alt={title}
            src={categorieAttribute}
            className=' object-cover h-full w-full group-hover:scale-110 transition'
            height={150}
            width={150}
          />
        </div>
        <div className='font-semibold text-lg '>{title} </div>
        <div className='font-semibold text-lg '>0/284 words </div>
        <div className='font-semibold text-lg '>infos sur le package </div>
      </div>
    </div>
  );
};
