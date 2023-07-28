import type { VercelRequest, VercelResponse } from "@vercel/node";
import axios from "axios";

export default async function handler(req: VercelRequest, res: VercelResponse) {
  try {
    const { apiUrl } = req.body;

    console.log("ehyyyyyy");
    const response = await axios.get(apiUrl);

    res.status(200).json(response.data);
  } catch (error: any) {
    console.error("Error in proxy request:", error.message);
    res
      .status(500)
      .json({ error: "Failed to fetch data from the external API" });
  }
}
