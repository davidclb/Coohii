type Props = {
  children: React.ReactNode;
};

export const FeedWrapper = ({ children }: Props) => {
  return <div className='relative top-0 pb-10'>{children}</div>;
};
