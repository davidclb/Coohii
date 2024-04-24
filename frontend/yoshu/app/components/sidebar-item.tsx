"use client";

import { ChevronDown } from "lucide-react";
import { Button } from "./ui/button";
import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/components/ui/accordion";

type Props = {
  label: string;
  iconSrc: string;
  href: string;
  items?: SubItemProps[];
};

type SubItemProps = {
  name: string;
  href: string;
};

export const SidebarItem = ({ label, iconSrc, href, items }: Props) => {
  const pathname = usePathname();
  const active = pathname === href;

  return (
    <>
      <Button
        className='justify-start h-[52px] '
        variant={active ? "sidebarOutline" : "sidebar"}
        asChild
      >
        <Link href={href}>
          <Image
            src={iconSrc}
            alt={label}
            className='mr-5'
            height={32}
            width={32}
          />
          {label}
        </Link>
      </Button>
      
    </>
  );
};