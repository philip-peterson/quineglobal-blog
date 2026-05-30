// Theme toggle with localStorage + system preference detection
(function () {
  const root = document.documentElement;
  const toggle = document.getElementById('theme-toggle');

  if (!toggle) return;

  const ICON_DARK = '<span class="icon">☾</span>'; // moon = go dark
  const ICON_LIGHT = '<span class="icon">☀︎</span>'; // sun = go light

  function applyTheme(theme) {
    if (theme === 'dark') {
      root.classList.add('dark');
      toggle.innerHTML = ICON_LIGHT;
      toggle.setAttribute('aria-label', 'Switch to light mode');
    } else {
      root.classList.remove('dark');
      toggle.innerHTML = ICON_DARK;
      toggle.setAttribute('aria-label', 'Switch to dark mode');
    }
  }

  function getPreferredTheme() {
    const saved = localStorage.getItem('theme');
    if (saved === 'dark' || saved === 'light') {
      return saved;
    }
    // Fall back to system preference
    return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
  }

  // Initialize
  const initialTheme = getPreferredTheme();
  applyTheme(initialTheme);

  // Click handler
  toggle.addEventListener('click', function () {
    const isCurrentlyDark = root.classList.contains('dark');
    const nextTheme = isCurrentlyDark ? 'light' : 'dark';
    applyTheme(nextTheme);
    localStorage.setItem('theme', nextTheme);
  });

  // Optional: react to system preference changes if user hasn't explicitly chosen
  const media = window.matchMedia('(prefers-color-scheme: dark)');
  media.addEventListener('change', function (e) {
    if (!localStorage.getItem('theme')) {
      applyTheme(e.matches ? 'dark' : 'light');
    }
  });
})();
