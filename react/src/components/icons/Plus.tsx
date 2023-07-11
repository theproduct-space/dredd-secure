export default function Plus(props: any): JSX.Element {
  return (
    <svg
      width="20"
      height="21"
      viewBox="0 0 20 21"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      {...props}
    >
      <path
        d="M9.75 2.25V18.75"
        stroke="#FAFAFA"
        stroke-width="3"
        stroke-linecap="round"
      />
      <path
        d="M1.5 10.5L18 10.5"
        stroke="#FAFAFA"
        stroke-width="3"
        stroke-linecap="round"
      />
    </svg>
  );
}
