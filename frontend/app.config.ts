export default defineAppConfig({
  ui: {
    primary: "mandy",
    gray: "cool",
    container: {
      constrained: "max-w-7xl",
    },
    button: {
      variant: {
        solid:
          "bg-mandy text-black hover:bg-mandy/90 text-lg rounded-sm px-6 py-2 uppercase",
        default:
          "bg-black border-lightRed text-lightRed hover:text-lightRed/90 border-2 hover:border-lightRed/50 text-lg rounded-sm px-6 py-2 uppercase",
        outline:
          "border border-2 border-lightRed bg-transparent hover:border-darkRed hover:text-lightRed text-lg rounded-sm px-6 py-2 uppercase",
        secondary:
          "bg-lightRed border-lightRed text-black hover:text-black/90 border-2 hover:border-medRed/50 hover:bg-medRed/90 text-lg rounded-sm px-6 py-2 uppercase",
        ghost: "hover:text-lightRed border-x-2 text-lg rounded-sm px-6 py-2",
        link: "text-primary underline-offset-4 hover:underline hover:text-lightRed text-lg rounded-sm px-6 py-2 uppercase",
      },
    },
    input: {
      variant: {
        default:
          "shadow-sm bg-black text-gray-900 dark:text-white ring-1 ring-inset ring-{color}-500 dark:ring-{color}-400 focus:ring-2 focus:ring-{color}-500 dark:focus:ring-{color}-400",
      },
      color: {
        black: {
          default:
            "shadow-sm bg-black text-white ring-1 ring-inset ring-gray-700 focus:ring-2 focus:ring-primary-400",
        },
      },
      default: {
        variant: "default",
        color: "black",
      },
    },
  },
});
