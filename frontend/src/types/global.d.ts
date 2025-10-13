declare global {
  interface Window {
    gtag?: (
      command: "config" | "event" | "js" | "set",
      targetId?: string,
      config?: {
        page_title?: string;
        page_location?: string;
        custom_map?: Record<string, string>;
        [key: string]: unknown;
      },
    ) => void;
  }
}

export {};
