"use client";

import Link from "next/link";
import { Button } from "./ui/button";
import { useAuth } from "../context/AuthContext";
import { useState } from "react";
import { LoginForm } from "./loginForm";

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
  // State related to Login Details of user

  const { isAuthenticated, user, login, logout } = useAuth();

  // State related to Login Form

  const [showModal, setShowModal] = useState(false);

  const handleOpenModal = () => {
    setShowModal(true);
  };

  const handleCloseModal = () => {
    setShowModal(false);
  };

  const handleAddReview = () => {
    if (isAuthenticated && user) {
      console.log("balance la requete");
    } else {
      handleOpenModal();
    }
  };

  return (
    <>
      <div className='py-1 w-full'>
        <div className='lg:flex items-center justify-center w-full'>
          <div className=' lg:mr-7 lg:mb-0 mb-7 bg-white w-full p-6 shadow rounded-xl'>
            <div className='flex align-center justify-center border-b border-gray-200 pb-6'>
              <p className='text-2xl block font-medium leading-5 text-gray-800'>
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
                  className={` bg-blue-500 focus:outline-none rounded-full text-white font-medium `}
                  onClick={handleAddReview}
                >
                  {" "}
                  Add to reviews
                </Button>
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
        <LoginForm
          show={showModal}
          onClose={handleCloseModal}
          onLogin={login}
        />
      </div>
    </>
  );
};
