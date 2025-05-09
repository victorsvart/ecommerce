type ProductCardProps = {
  id: number;
  name: string;
  description: string;
  price: number;
  discountPercentage: number | null;
  imageUrl: string;
  href: string;
};

export function ProductCard({
  id,
  name,
  description,
  imageUrl,
  price,
  discountPercentage,
  href,
}: ProductCardProps) {
  const productId = id;
  return (
    <div className="max-w-sm h-full flex flex-col bg-white border border-gray-200 rounded-lg shadow-sm dark:bg-gray-800 dark:border-gray-700">
      <a href={href}>
        <img
          className="w-full h-80 object-cover rounded-t-lg"
          src={imageUrl}
          alt={name}
        />
      </a>
      <div className="flex flex-col flex-grow p-5">
        <a href={href}>
          <p className="mb-2 font-bold tracking-tight text-gray-900 dark:text-white">
            {name}
          </p>
        </a>
        <p className="mb-3 text-xl font-bold text-rose-400">{price}</p>
        <div className="mt-auto flex justify-center">
          <a
            href={href}
            className="w-full flex items-center justify-center px-3 py-2 text-sm font-medium text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-rose-600 dark:hover:bg-rose-700 dark:focus:ring-blue-800 transition-all duration-300 ease-in-out"
          >
            <p>Adicionar ao carrinho</p>
            <svg
              className="rtl:rotate-180 w-3.5 h-3.5 ms-2"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 14 10"
            >
              <path
                stroke="currentColor"
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="M1 5h12m0 0L9 1m4 4L9 9"
              />
            </svg>
          </a>
        </div>
      </div>
    </div>
  );
}
