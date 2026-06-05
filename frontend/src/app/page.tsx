"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";

export default function Home() {
  const router = useRouter();
  const [url, setUrl] = useState("");
  const [shortUrl, setShortUrl] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [expiresIn, setExpiresIn] = useState<string>("");

useEffect(() => {
    const savedToken = localStorage.getItem("token");
    if (!savedToken) {
      router.push("/login");
    }
}, []);
const handleSubmit = async () => {
    if (!url) return;
    setLoading(true);
    setError("");
    setShortUrl("");

    const savedToken = localStorage.getItem("token");

    try {
      const res = await fetch("http://localhost:8080/api/urls", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${savedToken}`,
        },
        body: JSON.stringify({
        original_url: url,
        expires_in: expiresIn ? parseInt(expiresIn) : null,
}),
      });

      const data = await res.json();

      if (!res.ok) {
        setError(data.error);
        return;
      }

      setShortUrl(`http://localhost:8080/${data.short_code}`);
    } catch (_: unknown) {
      setError("サーバーに接続できません");
    } finally {
      setLoading(false);
    }
  };

  const handleLogout = () => {
    localStorage.removeItem("token");
    router.push("/login");
  };

  return (
    <main className="min-h-screen flex items-center justify-center bg-gray-50">
      <div className="bg-white p-8 rounded-xl shadow-md w-full max-w-md">
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-2xl font-bold">URL短縮</h1>
          <button
            onClick={handleLogout}
            className="text-sm text-gray-500 hover:text-red-500"
          >
            ログアウト
          </button>
        </div>

        <input
          type="text"
          placeholder="https://example.com"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          className="w-full border rounded-lg px-4 py-2 mb-4 focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <select
  value={expiresIn}
  onChange={(e) => setExpiresIn(e.target.value)}
  className="w-full border rounded-lg px-4 py-2 mb-4 focus:outline-none focus:ring-2 focus:ring-blue-500"
>
  <option value="">期限なし</option>
  <option value="1">1時間</option>
  <option value="24">24時間</option>
  <option value="168">7日間</option>
</select>
        <button
          onClick={handleSubmit}
          disabled={loading}
          className="w-full bg-blue-500 text-white py-2 rounded-lg hover:bg-blue-600 disabled:opacity-50"
        >
          {loading ? "処理中..." : "短縮する"}
        </button>

        {error && (
          <p className="mt-4 text-red-500 text-sm">{error}</p>
        )}

        {shortUrl && (
          <div className="mt-4 p-4 bg-gray-50 rounded-lg">
            <p className="text-sm text-gray-500 mb-1">短縮URL</p>
            <a
  href={shortUrl}
  target="_blank"
  className="text-blue-500 break-all hover:underline"
  >
  {shortUrl}
</a>
          </div>
        )}
      </div>
    </main>
  );
}