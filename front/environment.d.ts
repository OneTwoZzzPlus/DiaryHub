declare global {
  namespace NodeJS {
    interface ProcessEnv {
      NEXT_PUBLIC_ADDRESS_AUTH: string;
    }
  }
}

export {};