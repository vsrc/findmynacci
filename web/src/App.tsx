import React, { useState, useEffect } from "react";
import axios from "axios";
import { Toaster, toast } from "react-hot-toast";

const App = () => {
  const [count, setCount] = useState("0");

  const fetchNumber = async (direction: string) => {
    try {
      const response = await axios.get(`http://localhost:8080/${direction}`);
      if (response.status === 200 && response.data?.current) {
        setCount(response.data.current);
      } else {
        toast.error("Server responded, but the format is incorrect.");
      }
    } catch (error) {
      toast.error("There was an issue connecting to the server.");
    }
  };

  useEffect(() => {
    fetchNumber("current");
  }, []);

  return (
    <div className="flex flex-col h-screen bg-[#1A202C] text-white">
      <Toaster position="top-center" />
      <header className="text-4xl p-5 text-center font-bold">
        Find my (fibo)nacci
      </header>

      <main className="flex-grow flex flex-col items-center justify-center">
        <div className="text-9xl font-bold">{count}</div>
        <div className="mt-8">
          <button
            className="btn btn-lg btn-accent m-2 w-32"
            onClick={() => fetchNumber("previous")}
            disabled={count === "0"}
          >
            Previous
          </button>

          <button
            className="btn btn-lg btn-secondary m-2 w-32"
            onClick={() => fetchNumber("current")}
          >
            Current
          </button>

          <button
            className="btn btn-lg btn-primary m-2 w-32"
            onClick={() => fetchNumber("next")}
          >
            Next
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
