"use client";

import { LoginForm } from "@/app/components/loginForm";
import { useAuth } from "@/app/context/AuthContext";
import { MagnifyingGlassIcon } from "@heroicons/react/24/outline";
import { useSearchParams, usePathname, useRouter } from "next/navigation";
import { useState } from "react";

const SearchBar = () => {

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

  const searchParams = useSearchParams();
  const pathname = usePathname();
  const { replace } = useRouter();


  const handleSearch = (searchTerm: string) => {

    const params = new URLSearchParams(searchParams);
    if (searchTerm){
        params.set("query", searchTerm);
    }
    else {
        params.delete("query")
    }
    replace(`${pathname}?${params.toString()}`);

  }

  return (
    <div className=' relative flex flex-1 flex shrink-0 space-x-40'>
      <label htmlFor='search' className='sr-only'>
        Search
      </label>
      <input
        className='peer block w-1/2 rounded-md border border-gray-200 pv-[9px] pl-10 text-sm outline-2 placeholder:text-grey-500'
        placeholder='search'
        defaultValue={searchParams.get("query")?.toString()}
        onChange={(e) => {
          handleSearch(e.target.value);
        }}
      />
      <MagnifyingGlassIcon className='absolute left-3 top-1/2 h-[18px] w-[18px] -translate-y-1/2 text-grey-500 peer-focus:text-grey-900' />

      {isAuthenticated && user ? (
        <div className='flex items-center space-x-4'>
          <span> {user.username}</span>
          <button
            onClick={logout}
            className='bg-red-500 text-white py-2 px-4 rounded-md hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500'
          >
            Logout
          </button>
        </div>
      ) : (
        <div className=' flex bg-gray-100'>
          <button
            onClick={handleOpenModal}
            className='bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500'
          >
            Login/Signup
          </button>
          <LoginForm show={showModal} onClose={handleCloseModal} onLogin={login} />
        </div>
      )}
    </div>
  );
};

export default SearchBar;
