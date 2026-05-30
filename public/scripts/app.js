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

/* ============================================
   Page Slide Transitions on Internal Navigation
   ============================================ */

(function () {
  if (!document.startViewTransition) return;

  // Helper to perform navigation with direction-aware slide
  async function navigateWithTransition(href, direction = 'forward') {
    try {
      const html = await fetch(href, { headers: { 'X-Requested-With': 'fetch' } }).then(r => r.text());
      const newDoc = new DOMParser().parseFromString(html, 'text/html');

      // Add direction class for CSS to pick the right animations
      const htmlEl = document.documentElement;
      if (direction === 'back') {
        htmlEl.classList.add('back-transition');
      }

      const transition = document.startViewTransition(() => {
        document.body.innerHTML = newDoc.body.innerHTML;
        document.title = newDoc.title;

        // Only pushState on forward navigation (popstate already updated history)
        if (direction === 'forward') {
          history.pushState({}, '', href);
        }

        reinitThemeToggle();
      });

      // Clean up direction class after transition
      transition.finished.finally(() => {
        htmlEl.classList.remove('back-transition');
      });
    } catch (err) {
      window.location.href = href;
    }
  }

  // Intercept internal link clicks (forward navigation)
  document.addEventListener('click', (event) => {
    const link = event.target.closest('a');
    if (!link) return;

    const url = new URL(link.href, window.location.href);
    if (url.origin !== window.location.origin) return;

    if (link.hasAttribute('download') || link.target === '_blank' || link.hasAttribute('data-no-transition')) {
      return;
    }

    event.preventDefault();
    const href = url.pathname + url.search + url.hash;

    navigateWithTransition(href, 'forward');
  });

  // Handle browser back/forward buttons
  window.addEventListener('popstate', () => {
    // When using back/forward, we want the reverse slide direction
    navigateWithTransition(location.href, 'back');
  });

  // Re-attach the theme toggle after body content is replaced
  function reinitThemeToggle() {
    const root = document.documentElement;
    const toggle = document.getElementById('theme-toggle');
    if (!toggle) return;

    const ICON_DARK = '<span class="icon">☾</span>';
    const ICON_LIGHT = '<span class="icon">☀︎</span>';

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
      if (saved === 'dark' || saved === 'light') return saved;
      return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
    }

    const initialTheme = getPreferredTheme();
    applyTheme(initialTheme);

    toggle.onclick = () => {
      const isDark = root.classList.contains('dark');
      const next = isDark ? 'light' : 'dark';
      applyTheme(next);
      localStorage.setItem('theme', next);
    };
  }
})();
