export interface MessageResponse {
  message: string;
}

export const fetchData = async (): Promise<MessageResponse | undefined> => {
  try {
    const response = await fetch("/api/hello");
    if (!response.ok) throw new Error("Network error");
    return (await response.json()) as MessageResponse;
  } catch (err) {
    console.error(err);
  }
};
