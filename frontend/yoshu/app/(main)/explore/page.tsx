import { FeedWrapper } from "@/app/components/feed-wrapper";
import { StickyWrapper } from "../../components/sticky-wrapper";

import { Header } from "./header";
import { UserProgress } from "@/app/components/user-progress";
import { Button } from "@/app/components/ui/button";
import { WordList } from "@/app/components/wordlist";
import Image from "next/image";

const ExplorePage = () => {
  return (
    <div className='flex flex-row-reverse gap-[48px] px-6'>
      <StickyWrapper>
        <UserProgress />
      </StickyWrapper>
      <FeedWrapper>
        <Header title='Search a hanzi, a word or a sentence ' />
        <div className=' max-w-[912px] px-3 mx-auto  mb-20'>
          <Button
            type='button'
            className='text-gray-900 border  border-white hover:border-gray-200 dark:border-gray-900 dark:bg-gray-900 dark:hover:border-gray-700 bg-white focus:ring-4 focus:outline-none focus:ring-gray-300 rounded-full text-base font-medium px-5 py-2.5 text-center me-3 mb-3 dark:text-white dark:focus:ring-gray-800'
          >
            HSK
          </Button>
          <Button
            type='button'
            className='text-gray-900 border border-white hover:border-gray-200 dark:border-gray-900 dark:bg-gray-900 dark:hover:border-gray-700 bg-white focus:ring-4 focus:outline-none focus:ring-gray-300 rounded-full text-base font-medium px-5 py-2.5 text-center me-3 mb-3 dark:text-white dark:focus:ring-gray-800'
          >
            New HSK
          </Button>
          <Button
            type='button'
            className='text-gray-900 border border-white hover:border-gray-200 dark:border-gray-900 dark:bg-gray-900 dark:hover:border-gray-700 bg-white focus:ring-4 focus:outline-none focus:ring-gray-300 rounded-full text-base font-medium px-5 py-2.5 text-center me-3 mb-3 dark:text-white dark:focus:ring-gray-800'
          >
            TOCFL
          </Button>
          <Button
            type='button'
            className='text-gray-900 border border-white hover:border-gray-200 dark:border-gray-900 dark:bg-gray-900 dark:hover:border-gray-700 bg-white focus:ring-4 focus:outline-none focus:ring-gray-300 rounded-full text-base font-medium px-5 py-2.5 text-center me-3 mb-3 dark:text-white dark:focus:ring-gray-800'
          >
            Custom
          </Button>
          <Button
            type='button'
            className='text-gray-900 border border-white hover:border-gray-200 dark:border-gray-900 dark:bg-gray-900 dark:hover:border-gray-700 bg-white focus:ring-4 focus:outline-none focus:ring-gray-300 rounded-full text-base font-medium px-5 py-2.5 text-center me-3 mb-3 dark:text-white dark:focus:ring-gray-800'
          >
            All
          </Button>
        </div>
        <div className=' px-5 grid auto-cols-max md:grid-cols-4 gap-3	 '>
          <div className=' bg-emerald-500 	'>
            <div>
              <Image
                className='h-auto max-w-full rounded-lg'
                src='/hsk1.jpeg'
                height={100}
                width={100}
                alt='Dragon'
              />
            </div>
          </div>
          <div className=' bg-emerald-500'>
            <Image
              className='h-auto max-w-full rounded-lg'
              src='/hsk2.jpeg'
              height={100}
              width={100}
              alt='Dragon'
            />
          </div>
          <div className=' bg-emerald-500'>
            <Image
              className='h-auto max-w-full rounded-lg'
              src='/hsk3.jpeg'
              height={100}
              width={100}
              alt='Dragon'
            />
          </div>
          <div className=' bg-emerald-500'>
            <Image
              className='h-auto max-w-full rounded-lg'
              src='/hsk4.jpeg'
              height={100}
              width={100}
              alt='Dragon'
            />
          </div>
          <div className=' bg-emerald-500'>
            <Image
              className='h-auto max-w-full rounded-lg'
              src='/hsk5.jpeg'
              height={100}
              width={100}
              alt='Dragon'
            />
          </div>
          <div className=' bg-emerald-500'>
            <Image
              className='h-auto max-w-full rounded-lg'
              src='/hsk6.jpeg'
              height={100}
              width={100}
              alt='Dragon'
            />
          </div>
        </div>
      </FeedWrapper>
    </div>
  );
};

export default ExplorePage;
