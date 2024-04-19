"use client"


import Image from "next/image";
import { WordListCategorie } from "./wordlist-categorie";
import { usePathname, useSearchParams } from "next/navigation";


/* type Props = {
  title: any;
  userProgress: any;
  src: any;
}; */

export const wordcategories = [
  {
    label: 'HSK',

  },
  {
    label: 'New HSK',

  },
  {
    label: 'TOCFL',

  },
  {
    label: 'Custom',

  },
  {
    label: 'All',

  },
]

export const WordListCategories = (/* { title, userProgress, src }: Props */) => {
   const searchParams = useSearchParams();
   const pathName = usePathname()
   const category = searchParams.get('category')

   const isMainPage = (pathName === '/')
   
  return (
    <div className=' max-w-[912px] px-3 mx-auto  mb-20'>
      {wordcategories.map((item) => (
        <WordListCategorie label={item.label} />
      ))}
    </div>
  );
};
