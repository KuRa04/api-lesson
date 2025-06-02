document.getElementById('fetchBtn').addEventListener('click', async () => {
  const n = document.getElementById('numInput').value;
  const url = `/api/hello?n=${encodeURIComponent(n)}`; // エンドポイントを修正
  const start = performance.now();
  try {
    const res = await fetch(url);
    if (!res.ok) {
      const errText = await res.text();
      console.error('エラー:', errText);
      return;
    }
    const data = await res.json();
    const end = performance.now();
    console.log('配列:', data);
    console.log('レスポンスタイム:', (end - start).toFixed(2), 'ms');
  } catch (e) {
    console.error('通信エラー:', e);
  }
});