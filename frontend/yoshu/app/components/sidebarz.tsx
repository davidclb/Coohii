"use client";
import {
  LucideIcon,
  LayoutDashboard,
  BadgeDollarSign,
  CircleUserRound,
  Settings,
  WalletCards,
} from "lucide-react";
import SidebarItem from "./item";
import Link from "next/link";
import Image from "next/image";



interface ISidebarItem {
  name: string;
  path: string;
  icon: string;
  items?: ISubItem[];
}

interface ISubItem {
  name: string;
  path: string;
}

const items: ISidebarItem[] = [
  {
    name: "Home",
    path: "/",
    icon: '/home.svg',
  },
  {
    name: "Explore",
    path: "/explore",
    icon: '/explore.svg',
    items: [
        {
            name: "Radical",
            path: "/explore/radical",
          },
        {
          name: "Hanzi",
          path: "/explore/hanzi",
        },
        {
          name: "Words",
          path: "/explore/words",
        },
        {
          name: "Sentences",
          path: "/explore/sentences",
        },
      ],
  },
  {
    name: "Learn",
    path: "/learn",
    icon: '/learn.svg',
    items: [
        {
          name: "Review",
          path: "review",
        },
        {
          name: "Study hanzi ",
          path: "hanzi",
        },
        {
          name: "Study words",
          path: "words",
        },
      ],
  },
  {
    name: "Leaderboard",
    path: "/leaderboard",
    icon: '/leaderboard.svg',
  },
  {
    name: "Settings",
    path: "/settings",
    icon: '/settings.svg',
  },
];

const Sidebarz = () => {
  return (
    <div className="fixed top-0 left-0 h-screen w-64 bg-white shadow-lg z-10 p-4">
      <div className="flex flex-col space-y-10 w-full">
      <Link href='/explore'>
        <div className='pt-8 pl-4 pb-7 flex items-center gap-x-3'>
          <Image src='/dragona.svg' height={40} width={40} alt='Dragon' />
          <h1 className='text-2xl font-extrabold text-sky-600 tracking-wide'>
            Yoshu
          </h1>
        </div>
      </Link>
        <div className="flex flex-col space-y-2">
          {items.map((item, index) => (
            <SidebarItem key={index} item={item} />
          ))}
        </div>
      </div>
    </div>
  );
};

export default Sidebarz;