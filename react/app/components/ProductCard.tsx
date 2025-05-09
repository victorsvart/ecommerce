type ProductCardProps = {
  id: number;
  name: string;
  description: string;
  price: string;
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
  return (
    <div className="max-w-sm h-full flex flex-col bg-white border border-gray-200 rounded-2xl shadow-md dark:bg-gray-900 dark:border-gray-700 transition-transform hover:scale-[1.02] duration-300">
      <a href={href} className="overflow-hidden rounded-t-2xl">
        <img
          src={imageUrl}
          alt={name}
          className="w-full aspect-[4/3] object-cover hover:scale-105 transition-transform duration-300"
        />
      </a>
      <div className="flex flex-col flex-grow p-5 gap-3">
        <a href={href}>
          <h3 className="text-lg font-semibold text-gray-800 dark:text-white hover:underline">
            {name}
          </h3>
        </a>
        <p className="text-sm text-gray-600 dark:text-gray-400 line-clamp-2">
          {description}
        </p>

        <div className="flex flex-col gap-1 mt-auto">
          {discountPercentage ? (
            <div className="flex items-center gap-2">
              <span className="text-lg font-bold text-rose-500">{price}</span>
              <span className="text-xs font-medium text-green-500">
                -{discountPercentage}%
              </span>
            </div>
          ) : (
            <span className="text-lg font-bold text-rose-500">{price}</span>
          )}
        </div>

        <a
          href={href}
          className="mt-3 w-full inline-flex items-center justify-center px-4 py-2 text-sm font-medium text-white bg-blue-600 rounded-xl hover:bg-blue-700 focus:ring-2 focus:outline-none focus:ring-blue-300 dark:bg-rose-600 dark:hover:bg-rose-700 dark:focus:ring-rose-500 transition-all"
        >
          <p>Adicionar ao carrinho</p>
          <svg
            className="w-4 h-4 ml-2"
            fill="none"
            stroke="currentColor"
            strokeWidth={2}
            viewBox="0 0 24 24"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              d="M17 8l4 4m0 0l-4 4m4-4H3"
            />
          </svg>
        </a>
      </div>
    </div>
  );
}
