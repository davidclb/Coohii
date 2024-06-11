import { Sidebar } from "../components/sidebar";
import { MobileHeader } from "../components/mobile-header";
import Sidebarz from "../components/sidebarz";

type Props = {
  children: React.ReactNode;
};

const MainLayout = ({ children }: Props) => {


  return (
    <>
      <MobileHeader />
     {/* <Sidebar className='hidden lg:flex' />  */}
      <Sidebarz />
      <main className='lg:pl-[256px] h-full pt-[50px] lg:pt-0'>
        <div className='max-w-[1056px] mx-auto pt-6 h-full'>{children}</div>
      </main>
    </>
  );
};

export default MainLayout;
function useState(arg0: boolean): [any, any] {
  throw new Error("Function not implemented.");
}

