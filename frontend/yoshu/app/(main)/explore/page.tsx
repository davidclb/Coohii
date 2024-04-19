import { FeedWrapper } from "@/app/components/feed-wrapper";
import { StickyWrapper } from "../../components/sticky-wrapper";
import { Header } from "./header";
import { UserProgress } from "@/app/components/user-progress";
import Image from "next/image";
import { WordListCategorie } from "@/app/components/wordlist-categorie";
import { WordListCategories } from "@/app/components/wordlist-categories";


async function getUser(){
        const res = await fetch(`http://http://127.0.0.1:3001/user`);
        return res.json();
}

const ExplorePage = () => {

  
  return (
    <div className='flex flex-row-reverse gap-[48px] px-6'>
      <StickyWrapper>
        <UserProgress />
      </StickyWrapper>
      <FeedWrapper>
        <Header title='Search a hanzi, a word or a sentence ' />
        <div className=' max-w-[912px] px-3 mx-auto  mb-20'>
          {/*           <WordListCategorie label='HSK' />
          <WordListCategorie label='New HSK' />
          <WordListCategorie label='TOCFL' />
          <WordListCategorie label='Custom' />
          <WordListCategorie label='All' /> */}
          <WordListCategories />
        </div>
        <div
          className='          
             bg-emerald-500
              pt-24
              grid
              grid-cols-1
              sm:grid-cols-2
              md:grid-cols-3
              lg:grid-cols-4
              xl:grid-cols-5
              2xl:grid-cols-6
              gap-8 '
        >
          {/*           <div className=' bg-emerald-500 	'>
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
          </div> */}

          <div className=' bg-blue-500'> My future listings</div>
          <div className=' bg-blue-500'> My future listings</div>
          <div className=' bg-blue-500'> My future listings</div>
          <div className=' bg-blue-500'> My future listings</div>
          <div className=' bg-blue-500'> My future listings</div>
          <div className=' bg-blue-500'> My future listings</div>
          <div className=' bg-blue-500'> My future listings</div>
          <div className=' bg-blue-500'> My future listings</div>
          <div className=' bg-blue-500'> My future listings</div>
          <div className=' bg-blue-500'> My future listings</div>
          <div className=' bg-blue-500'> My future listings</div>
          <div className=' bg-blue-500'> My future listings</div>
          <div className=' bg-blue-500'> My future listings</div>
          <div className=' bg-blue-500'> My future listings</div>
          <div className=' bg-blue-500'> My future listings</div>
          <div className=' bg-blue-500'> My future listings</div>
          <div className=' bg-blue-500'> My future listings</div>
          <div className=' bg-blue-500'> My future listings</div>
        </div>
      </FeedWrapper>
    </div>
  );
};

export default ExplorePage;
