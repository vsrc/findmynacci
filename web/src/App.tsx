import React, { useState } from "react";

const App = () => {
  const [count, setCount] = useState(0);

  const increment = () => setCount((prevCount) => prevCount + 1);
  const decrement = () =>
    setCount((prevCount) => (prevCount > 0 ? prevCount - 1 : 0));

  return (
    <div className="flex flex-col h-screen bg-[#1A202C] text-white">
      <header className="text-4xl p-5 text-center font-bold">
        Find my (fibo)nacci
      </header>

      <main className="flex-grow flex flex-col items-center justify-center">
        <div className="text-9xl font-bold">{count}</div>
        <div className="mt-8">
          <button
            className="btn btn-lg btn-accent m-2"
            onClick={decrement}
            disabled={count === 0}
          >
            Decrease
          </button>
          <button className="btn btn-lg btn-primary m-2" onClick={increment}>
            Increase
          </button>
        </div>
      </main>

      <footer className="text-center p-5">
        <a
          target="_blank"
          rel="noreferrer"
          href="https://github.com/vsrc?tab=repositories"
          className="link link-hover"
        >
          github.com/vsrc
        </a>
      </footer>
    </div>
  );
};

export default App;
