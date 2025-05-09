type ProductCardProps = {
  id: number;
  name: string;
  description: string;
  imageUrl: string;
  href: string;
};

export function ProductCard({
  id,
  name,
  description,
  imageUrl,
  href,
}: ProductCardProps) {
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
          <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
            {name}
          </h5>
        </a>
        <p className="mb-3 font-normal text-gray-700 dark:text-gray-400">
          {description}
        </p>
        <div className="mt-auto flex justify-end">
          <a
            href={href}
            className="inline-flex items-center px-3 py-2 text-sm font-medium text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-orange-600 dark:hover:bg-orange-700 dark:focus:ring-blue-800 transition-all duration-300 ease-in-out"
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
