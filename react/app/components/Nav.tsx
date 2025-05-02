export const Navbar = () => {
  return (
    <nav className="flex w-full sticky shadow-md p-4 dark:bg-gray-900">
      <div className="container mx-auto flex items-center justify-between">
        <div className="flex items-center">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="20"
            height="20"
            fill="currentColor"
            className="text-gray-500"
            viewBox="0 0 16 16"
          >
            <path d="m.334 0 4.358 4.359h7.15v7.15l4.358 4.358V0zM.2 9.72l4.487-4.488v6.281h6.28L6.48 16H.2z" />
          </svg>
          <span className="ml-2 text-lg font-semibold text-gray-700 invisible sm:visible">
            ECommerce
          </span>
        </div>

        <div className="absolute left-1/2 transform -translate-x-1/2 flex space-x-6">
          <a href="#" className="text-gray-600 hover:text-gray-900">
            <p>Início</p>
          </a>
          <a href="#" className="text-gray-600 hover:text-gray-900">
            <p>Pesquise</p>
          </a>
          <a href="#" className="text-gray-600 hover:text-gray-900">
            <p>Contato</p>
          </a>
        </div>

        <div className="flex items-center space-x-4">
          <div className="flex justify-center items-center w-7 h-7 rounded-full bg-blue-700">
            V
          </div>
        </div>
      </div>
    </nav>
  );
};
