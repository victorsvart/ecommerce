interface AlertErrorProps {
  message: string | null;
  onClose: () => void;
}

export const AlertError: React.FC<AlertErrorProps> = ({ message, onClose }) => {
  if (!message) return null;

  return (
    <div
      className="flex items-start p-4 mb-4 text-sm text-red-800 border border-red-300 rounded-lg bg-red-50 dark:bg-gray-800 dark:text-red-400 dark:border-red-800"
      role="alert"
    >
      <svg
        className="shrink-0 w-4 h-4 me-3 mt-1"
        aria-hidden="true"
        xmlns="http://www.w3.org/2000/svg"
        fill="currentColor"
        viewBox="0 0 20 20"
      >
        <path d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5ZM9.5 4a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM12 15H8a1 1 0 0 1 0-2h1v-3H8a1 1 0 0 1 0-2h2a1 1 0 0 1 1 1v4h1a1 1 0 0 1 0 2Z" />
      </svg>
      <div className="flex-grow">
        <span className="font-medium">{message}</span>
      </div>
      <button
        onClick={onClose}
        className="ml-4 text-red-800 hover:text-red-600 dark:text-red-400 dark:hover:text-red-200"
        aria-label="Close alert"
      >
        ×
      </button>
    </div>
  );
};
