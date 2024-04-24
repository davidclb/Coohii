import { FeedWrapper } from "@/app/components/feed-wrapper";
import { StickyWrapper } from "../../../components/sticky-wrapper";
import { Header } from "../header";
import { UserProgress } from "@/app/components/user-progress";
import { WordListCategories } from "@/app/components/wordlist-categories";
import { CategorieCard } from "@/app/components/categorie-card";
import { usePathname, useSearchParams } from "next/navigation";

export const categories = [
  {
    label: "HSK1",
    type: "HSK",
    image: "/hsk1.jpeg",
  },
  {
    label: "HSK2",
    type: "HSK",
    image: "/hsk2.jpeg",
  },
  {
    label: "HSK3",
    type: "HSK",
    image: "/hsk3.jpeg",
  },
  {
    label: "HSK4",
    type: "HSK",
    image: "/hsk4.jpeg",
  },
  {
    label: "HSK5",
    type: "HSK",
    image: "/hsk5.jpeg",
  },
  {
    label: "HSK6",
    type: "HSK",
    image: "/hsk1.jpeg",
  },
  {
    label: "New HSK1",
    type: "New HSK",
    image: "/hsk5.jpeg",
  },
  {
    label: "New HSK2",
    type: "New HSK",
    image: "/hsk2.jpeg",
  },
  {
    label: "New HSK3",
    type: "New HSK",
    image: "/hsk3.jpeg",
  },
  {
    label: "New HSK4",
    type: "New HSK",
    image: "/hsk4.jpeg",
  },
  {
    label: "New HSK5",
    type: "New HSK",
    image: "/hsk5.jpeg",
  },
  {
    label: "HSK1",
    type: "HSK",
    image: "/hsk1.jpeg",
  },
  {
    label: "HSK2",
    type: "HSK",
    image: "/hsk2.jpeg",
  },
  {
    label: "HSK3",
    type: "HSK",
    image: "/hsk3.jpeg",
  },
  {
    label: "HSK4",
    type: "HSK",
    image: "/hsk4.jpeg",
  },
  {
    label: "HSK5",
    type: "HSK",
    image: "/hsk5.jpeg",
  },
  {
    label: "HSK6",
    type: "HSK",
    image: "/hsk1.jpeg",
  },
  {
    label: "New HSK1",
    type: "New HSK",
    image: "/hsk5.jpeg",
  },
  {
    label: "New HSK2",
    type: "New HSK",
    image: "/hsk2.jpeg",
  },
  {
    label: "New HSK3",
    type: "New HSK",
    image: "/hsk3.jpeg",
  },
  {
    label: "New HSK4",
    type: "New HSK",
    image: "/hsk4.jpeg",
  },
  {
    label: "New HSK5",
    type: "New HSK",
    image: "/hsk5.jpeg",
  },
];

/* async function getUser() {
  const res = await fetch(`http://localhost:3001/user`);
  return res.json();
} */

export default function ExplorePage() {
  /* const userData = await getUser(); */
  
  return (
    <div className='flex flex-row-reverse gap-[48px] px-6'>
      <StickyWrapper>
        <UserProgress />
      </StickyWrapper>
      <FeedWrapper>
        <Header title='Search a hanzi, a word or a sentence ' />
        <div className=' max-w-[912px] px-3 mx-auto  mb-20'>
          <WordListCategories />
        </div>
        <div
          className='          
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

          {categories.map((categorie: any) => {
            return (
              <CategorieCard
                title={categorie.label}
                categorieAttribute={categorie.image}
              />
            );
          })}
        </div>
      </FeedWrapper>
    </div>
  );
}
