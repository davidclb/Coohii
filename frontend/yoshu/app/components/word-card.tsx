import Link from "next/link";
import { Button } from "./ui/button";

type Props = {
  category: string;
  type: string;
  image: string;
  carac_simpl: string;
  carac_trad: string;
  pinyin: string;
  meaning: string;
};

export const WordCard = ({
  category,
  type,
  image,
  carac_simpl,
  carac_trad,
  pinyin,
  meaning,
}: Props) => {
  return (
    <>
      <div className='py-1 w-full'>
        <div className='lg:flex items-center justify-center w-full'>
          <div className=' lg:mr-7 lg:mb-0 mb-7 bg-white w-full p-6 shadow rounded-xl'>
            <div className='flex items-center border-b border-gray-200 pb-6'>
              <p className='text-2xl font-medium leading-5 text-gray-800'>
                <Link href=' /search'>
                  <Button>
                    {carac_simpl} ({carac_simpl})
                  </Button>
                </Link>
              </p>
              <div className='flex items-start justify-between w-full'>
                <div className='pl-3 w-full'>
                  <p className='text-xl font-medium leading-5 text-gray-800'>
                    ( {carac_trad} )
                  </p>
                </div>
                <Button
                  type='button'
                  className={` 'text-gray-900 border  border-blue-500 hover:border-gray-200 dark:border-gray-900 dark:bg-gray-900 dark:hover:border-gray-700 bg-blue-500 focus:ring-4 focus:outline-none focus:ring-gray-300 rounded-full text-base font-medium px-5 py-2.5 text-center me-3 mb-3 dark:text-white dark:focus:ring-gray-800`}
                ></Button>
              </div>
            </div>
            <div className='px-2'>
              <p className='text-sm leading-5 py-4 text-gray-600'>{meaning}</p>
              <div className='flex'>
                <div className='py-2 px-4 text-xs font-bold leading-3 text-white rounded-full bg-green-400'>
                  {category}
                </div>
                <div className='py-2 px-4 ml-3 text-xs font-bold leading-3 text-white rounded-full bg-red-400'>
                  RTH2
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};
