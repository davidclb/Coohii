import { cn } from "@/lib/utils";
import Image from "next/image";
import Link from "next/link";
import { SidebarItem } from "./sidebar-item";

type Props = {
  className?: string;
};

export const Sidebar = ({ className }: Props) => {
  return (
    <div
      className={cn(
        "flex h-full lg:w-[256px] lg:fixed left-0 top-0 px-4 border-r-2 flex-col",
        className
      )}
    >
      <Link href='/home'>
        <div className='pt-8 pl-4 pb-7 flex items-center gap-x-3'>
          <Image src='/dragona.svg' height={40} width={40} alt='Dragon' />
          <h1 className='text-2xl font-extrabold text-sky-600 tracking-wide'>
            Yoshu
          </h1>
        </div>
      </Link>
      <div className=' flex flex-col gap-y-2 flex-1'>
        <SidebarItem label='Home' href='/' iconSrc='/home.svg' />
        <SidebarItem label='Explore' href='/explore' iconSrc='/explore.svg' />
        <SidebarItem label='Learn' href='/learn' iconSrc='/learn.svg' />
        <SidebarItem
          label='Leaderboard'
          href='/leaderboard'
          iconSrc='/leaderboard.svg'
        />
        <SidebarItem
          label='Settings'
          href='/settings'
          iconSrc='/settings.svg'
        />
      </div>
    </div>
  );
};
