import { Login } from "./loginForm";
import { Button } from "./ui/button";
import React, { useState } from "react";

type Props = {
  activeCourse: any;
  hasActiveSubscription: boolean;
};



export const UserProgress = () => {
  return (
        <div>
      <Button className='bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-full'>
        Button
      </Button>
    </div> 
   
  );
};
