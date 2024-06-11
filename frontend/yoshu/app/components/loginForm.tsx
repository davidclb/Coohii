"use client";

import { FormEvent } from "react";

type LoginFormProps = {
  show: boolean;
  onClose: () => void;
  onLogin: (username: string) => void;
};

export const LoginForm: React.FC<LoginFormProps> = ({
  show,
  onClose,
  onLogin,
}) => {
  if (!show) return null;

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    const formData = new FormData(event.currentTarget);
    const username = formData.get("username");
    const password = formData.get("password");

    console.log(JSON.stringify({ username, password }));


    const response = await fetch("http://localhost:3001/signin", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ username, password }),
      
    });

    if (response.ok) {
      console.log("connectay");
      //router.push("/profile");

      if (username != null) {
        const usernamestring = username.toString();
        onLogin(usernamestring); ////////////////////////////////////// A MODIFIER //////////////////////////////
      }
    } else {
      // Handle errors
      console.log("pas bon");
    }
  }

  return (
    <>
      <div
        className='fixed inset-0 bg-black bg-opacity-70 z-50'
        onClick={onClose}
      ></div>
      <div className='fixed inset-0 flex items-center justify-center z-50'>
        <div className='bg-white p-8 rounded-lg shadow-lg relative'>
          <h2 className='text-2xl mb-4'>Formulaire</h2>
          <form onSubmit={handleSubmit}>
            <div className='mb-4'>
              <label className='block text-sm font-medium text-gray-700'>
                Username:
              </label>
              <input
                type='username'
                name='username'
                required
                className='mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm'
              />
            </div>
            <div className='mb-4'>
              <label className='block text-sm font-medium text-gray-700'>
                Mot de passe:
              </label>
              <input
                type='password'
                name='password'
                required
                className='mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm'
              />
            </div>
            <button
              type='submit'
              className='w-full bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500'
            >
              Soumettre
            </button>
          </form>
          <button
            onClick={onClose}
            className='absolute top-2 right-2 text-gray-500 hover:text-gray-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500'
          >
            <svg
              xmlns='http://www.w3.org/2000/svg'
              className='h-6 w-6'
              fill='none'
              viewBox='0 0 24 24'
              stroke='currentColor'
              strokeWidth='2'
            >
              <path
                strokeLinecap='round'
                strokeLinejoin='round'
                d='M6 18L18 6M6 6l12 12'
              />
            </svg>
          </button>
        </div>
      </div>
    </>
  );
};

/*     <form onSubmit={handleSubmit}>
      <input type='email' name='email' placeholder='Email' required />
      <input type='password' name='password' placeholder='Password' required />
      <button type='submit'>Login</button>
    </form> */
