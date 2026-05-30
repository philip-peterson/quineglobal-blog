// Theme toggle + Other Posts re-shuffle (using jQuery for simplicity)
$(function () {
  // ========== Theme Toggle ==========
  const $root = $(document.documentElement);
  const $toggle = $('#theme-toggle');

  if ($toggle.length === 0) return;

  const ICON_DARK = '<span class="icon">☾</span>';
  const ICON_LIGHT = '<span class="icon">☀︎</span>';

  function applyTheme(theme) {
    if (theme === 'dark') {
      $root.addClass('dark');
      $toggle.html(ICON_LIGHT);
      $toggle.attr('aria-label', 'Switch to light mode');
    } else {
      $root.removeClass('dark');
      $toggle.html(ICON_DARK);
      $toggle.attr('aria-label', 'Switch to dark mode');
    }
  }

  function getPreferredTheme() {
    const saved = localStorage.getItem('theme');
    if (saved === 'dark' || saved === 'light') {
      return saved;
    }
    return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
  }

  const initialTheme = getPreferredTheme();
  applyTheme(initialTheme);

  $toggle.on('click', function () {
    const isDark = $root.hasClass('dark');
    const next = isDark ? 'light' : 'dark';
    applyTheme(next);
    localStorage.setItem('theme', next);
  });

  // System preference changes
  const media = window.matchMedia('(prefers-color-scheme: dark)');
  media.addEventListener('change', function (e) {
    if (!localStorage.getItem('theme')) {
      applyTheme(e.matches ? 'dark' : 'light');
    }
  });

  // ========== Other Posts Re-shuffle (with heavy logging) ==========
  function initReshuffle() {
    console.log('[Re-shuffle] initReshuffle() called');

    const $section = $('.other-posts');
    const $container = $('.scattered-posts');
    const $btn = $('#reshuffle-posts');

    console.log('[Re-shuffle] Elements found:', {
      section: $section.length,
      container: $container.length,
      button: $btn.length
    });

    if ($section.length === 0 || $container.length === 0 || $btn.length === 0) {
      console.warn('[Re-shuffle] Missing required elements. Aborting.');
      return;
    }

    let allPosts = [];
    try {
      const b64 = $section.attr('data-other-posts') || '';
      const raw = b64 ? atob(b64) : '[]';
      allPosts = JSON.parse(raw);
      console.log('[Re-shuffle] Parsed allPosts from data attr, count =', allPosts.length);
      if (allPosts.length > 0) {
        console.log('[Re-shuffle] First post sample:', allPosts[0]);
      }
    } catch (e) {
      console.error('[Re-shuffle] Failed to parse data-other-posts JSON', e);
      return;
    }

    // Client-side layout generator (matches spirit of Go bestScatteredLayout)
    function generateLayout(n) {
      const layout = [];
      for (let i = 0; i < n; i++) {
        layout.push({
          top: Math.random() * 66 + 5,   // 5%–71%  (stays inside the ~220px container)
          left: Math.random() * 74 + 4,  // 4%–78%
          rot: Math.random() * 14 - 7,   // -7° to +7°
        });
      }
      return layout;
    }

    function scoreLayout(layout) {
      if (layout.length <= 1) return 999;
      let minDist = 1000;
      for (let i = 0; i < layout.length; i++) {
        for (let j = i + 1; j < layout.length; j++) {
          const dx = layout[i].left - layout[j].left;
          const dy = layout[i].top - layout[j].top;
          const dist = Math.sqrt(dx * dx + dy * dy);
          if (dist < minDist) minDist = dist;
        }
      }
      return minDist;
    }

    function bestScatteredLayout(n) {
      if (n <= 0) return [];
      const attempts = 70;
      let best = generateLayout(n);
      let bestScore = scoreLayout(best);
      for (let i = 0; i < attempts; i++) {
        const candidate = generateLayout(n);
        const s = scoreLayout(candidate);
        if (s > bestScore) {
          bestScore = s;
          best = candidate;
        }
      }
      return best;
    }

    function shuffle(arr) {
      console.log('[Re-shuffle] Shuffling array of length', arr.length);
      const a = arr.slice();
      for (let i = a.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [a[i], a[j]] = [a[j], a[i]];
      }
      return a;
    }

    $section.on('click', '#reshuffle-posts', function (e) {
      console.log('[Re-shuffle] Button clicked!');
      e.preventDefault();

      if (allPosts.length === 0) {
        console.warn('[Re-shuffle] Clicked but allPosts is empty');
        return;
      }

      const shuffled = shuffle(allPosts);
      const selected = shuffled.slice(0, Math.min(4, shuffled.length));
      console.log('[Re-shuffle] Selected posts for re-render:', selected.map(p => p.title));

      // Fresh varied layout every reshuffle (no more static 5 positions)
      const layout = bestScatteredLayout(selected.length);
      console.log('[Re-shuffle] Generated fresh layout:', layout);

      console.log('[Re-shuffle] Clearing .scattered-posts container');
      $container.empty();

      selected.forEach((post, i) => {
        const pos = layout[i];
        console.log(`[Re-shuffle] Creating card ${i}:`, { title: post.title, pos });

        const $a = $('<a>')
          .addClass('scattered-post')
          .attr('href', '/post/' + post.id)
          .css({
            top: pos.top + '%',
            left: pos.left + '%',
            transform: 'rotate(' + pos.rot + 'deg)'
          })
          .text(post.title);

        $container.append($a);
      });

      console.log('[Re-shuffle] Finished rendering', selected.length, 'cards');
    });

    console.log('[Re-shuffle] initReshuffle complete. Waiting for button click...');
  }

  initReshuffle();
});
