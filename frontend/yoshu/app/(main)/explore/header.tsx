import { Button } from "@/app/components/ui/button";
import { ArrowLeft } from "lucide-react";
import Link from "next/link";

type Props = {
  title: string;
};

export const Header = ({ title }: Props) => {
  return (
    <div className='sticky top-0 bg-white pb-3 lg:pt-[28px] flex items-center justify-between border-b-2 mb-5 text-neutral-400 lg:z-50'>
      <Link href='/search '>
        <Button>
          <ArrowLeft className='h-5 w-5 stroke-2 text-neutral-400' />
        </Button>
      </Link>
      <h1 className='font-bold text-lg'>{title}</h1>
      <div />
      <Button
        type='button'
        className={` 'text-gray-900 border  border-blue-500 hover:border-gray-200 dark:border-gray-900 dark:bg-gray-900 dark:hover:border-gray-700 bg-blue-500 focus:ring-4 focus:outline-none focus:ring-gray-300 rounded-full text-base font-medium px-5 py-2.5 text-center me-3 mb-3 dark:text-white dark:focus:ring-gray-800`}
      ></Button>
    </div>
  );
};
