import aiohttp
import os
import asyncio # <--- Need this to run the coroutine

GIPHY_TOKEN = os.getenv("GIPHY_TOKEN")

async def fetch_gif(tag: str) -> str | None:
    url = "https://api.giphy.com/v1/gifs/random"
    params = {
        "api_key": GIPHY_TOKEN,
        "tag": tag,
        "rating": "g",
    }

    async with aiohttp.ClientSession() as session:
        async with session.get(url, params=params) as response:
            if response.status != 200:
                print(f"Error: {response.status}")
                return None

            data = await response.json()
            gif_data = data.get("data", {})

            # Prefer MP4
            mp4 = gif_data.get("images", {}).get("original_mp4", {}).get("mp4")
            if mp4:
                return mp4

            # Fallback to GIF
            return gif_data.get("images", {}).get("original", {}).get("url")

# --- ADD THIS PART TO RUN IT ---
async def main():
    tag = "sunday" # Example tag
    result = await fetch_gif(tag)
    if result:
        print(f"Success! URL: {result}")
    else:
        print("Failed to fetch GIF.")

if __name__ == "__main__":
    asyncio.run(main())