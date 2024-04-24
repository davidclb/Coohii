import { Menu } from "lucide-react";

import { Sheet, SheetContent, SheetTrigger } from "./ui/sheet";

import { Sidebar } from "./sidebar";
import Sidebarz from "./sidebarz";

export const MobileSidebar = () => {
  return (
    <Sheet>
      <SheetTrigger>
        <Menu className='text-white' />
      </SheetTrigger>
      <SheetContent className='p-0 z-[100]' side='left'>
        <Sidebarz />
      </SheetContent>
    </Sheet>
  );
};
