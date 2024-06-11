"use client"
import SearchBar from "../explore/searchBar";
import { useState } from "react";


/* async function getUser() {
  const res = await fetch(`http://localhost:3001/user`);
  return res.json();
} */

export default function HomePage() {
  /* const userData = await getUser(); */


  return (
    <>
      <div className='flex gap-[48px] px-6'>
        <SearchBar />
      </div>
     
    </>
  );
}


