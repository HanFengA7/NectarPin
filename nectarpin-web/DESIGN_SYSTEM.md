# NectarPin Admin Design System Specification

## 1. Overview

### 1.1 Product Description
NectarPin Admin is a content management dashboard for a publishing platform, providing centralized management of articles, comments, and system configuration.

### 1.2 Design Philosophy
- **Modern & Professional**: Clean lines, subtle shadows, consistent spacing
- **Accessibility First**: WCAG AA compliant, keyboard navigation support
- **Responsive**: Adaptive layout across all device sizes
- **Dark Mode Support**: Full dark mode with prefers-color-scheme

### 1.3 Target Users
- Content administrators and editors
- Site moderators and reviewers
- System configuration managers

---

## 2. Design Tokens

### 2.1 Color System

#### Primary Colors
| Token | Light Mode | Dark Mode | Usage |
|-------|-----------|-----------|-------|
| `--color-primary` | #3B82F6 | #60A5FA | Primary actions, links |
| `--color-primary-hover` | #2563EB | #3B82F6 | Hover states |
| `--color-primary-active` | #1D4ED8 | #2563EB | Active/pressed states |
| `--color-primary-soft` | #EFF6FF | #1E3A5F | Soft backgrounds |
| `--color-primary-strong` | #1E40AF | #93C5FD | Strong emphasis |

#### Secondary Colors
| Token | Light Mode | Dark Mode | Usage |
|-------|-----------|-----------|-------|
| `--color-secondary` | #6366F1 | #818CF8 | Secondary actions |
| `--color-accent` | #8B5CF6 | #A78BFA | Accent highlights |

#### Semantic Colors
| Token | Light Mode | Dark Mode | Usage |
|-------|-----------|-----------|-------|
| `--color-success` | #10B981 | #34D399 | Success states |
| `--color-warning` | #F59E0B | #FBBF24 | Warning states |
| `--color-danger` | #EF4444 | #F87171 | Error/danger states |

#### Surface Colors
| Token | Light Mode | Dark Mode | Usage |
|-------|-----------|-----------|-------|
| `--color-bg` | #F8FAFC | #0F172A | Page background |
| `--color-surface` | #FFFFFF | #1E293B | Card surfaces |
| `--color-surface-alt` | #F1F5F9 | #334155 | Alternate surfaces |
| `--color-surface-hover` | #E2E8F0 | #475569 | Hover states |

#### Text Colors
| Token | Light Mode | Dark Mode | Usage |
|-------|-----------|-----------|-------|
| `--color-text` | #0F172A | #F1F5F9 | Primary text |
| `--color-text-primary` | #0F172A | #F1F5F9 | Headlines |
| `--color-text-secondary` | #475569 | #CBD5E1 | Body text |
| `--color-text-muted` | #64748B | #94A3B8 | Captions |
| `--color-text-placeholder` | #94A3B8 | #64748B | Placeholder text |
| `--color-text-disabled` | #CBD5E1 | #475569 | Disabled text |

#### Border Colors
| Token | Light Mode | Dark Mode | Usage |
|-------|-----------|-----------|-------|
| `--color-border` | #E2E8F0 | #334155 | Default borders |
| `--color-border-hover` | #CBD5E1 | #475569 | Hover borders |
| `--color-border-focus` | #3B82F6 | #60A5FA | Focus rings |

---

### 2.2 Typography

#### Font Families
```css
--font-sans: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
--font-heading: 'Plus Jakarta Sans', var(--font-sans);
--font-mono: 'JetBrains Mono', 'Fira Code', Consolas, monospace;
```

#### Font Sizes
| Token | Size | Usage |
|-------|------|-------|
| `--font-size-xs` | 0.75rem | Captions, badges |
| `--font-size-sm` | 0.8125rem | Secondary text |
| `--font-size-base` | 0.875rem | Body text |
| `--font-size-md` | 0.9375rem | Large body |
| `--font-size-lg` | 1rem | Small headings |
| `--font-size-xl` | 1.125rem | Section titles |
| `--font-size-2xl` | 1.25rem | Card titles |
| `--font-size-3xl` | 1.5rem | Page titles |
| `--font-size-4xl` | 1.875rem | Hero headings |
| `--font-size-5xl` | 2.25rem | Display text |

#### Font Weights
| Token | Value | Usage |
|-------|-------|-------|
| `--font-weight-normal` | 400 | Body text |
| `--font-weight-medium` | 500 | Labels |
| `--font-weight-semibold` | 600 | Buttons, nav |
| `--font-weight-bold` | 700 | Headings |
| `--font-weight-extrabold` | 800 | Hero text |

#### Line Heights
| Token | Value | Usage |
|-------|-------|-------|
| `--line-height-tight` | 1.25 | Headlines |
| `--line-height-snug` | 1.375 | Subheadings |
| `--line-height-normal` | 1.5 | Body text |
| `--line-height-relaxed` | 1.625 | Long text |
| `--line-height-loose` | 1.75 | Captions |

#### Letter Spacing
| Token | Value | Usage |
|-------|-------|-------|
| `--letter-spacing-tight` | -0.025em | Headlines |
| `--letter-spacing-normal` | 0 | Body |
| `--letter-spacing-wide` | 0.025em | Labels |
| `--letter-spacing-wider` | 0.05em | Caps |

---

### 2.3 Spacing System

#### Base Unit: 4px

| Token | Value | Usage |
|-------|-------|-------|
| `--space-0` | 0px | None |
| `--space-1` | 4px | Tight |
| `--space-2` | 8px | Small |
| `--space-3` | 12px | Medium-small |
| `--space-4` | 16px | Medium |
| `--space-5` | 20px | Medium-large |
| `--space-6` | 24px | Large |
| `--space-8` | 32px | XLarge |
| `--space-10` | 40px | 2XLarge |
| `--space-12` | 48px | 3XLarge |
| `--space-16` | 64px | 4XLarge |
| `--space-20` | 80px | 5XLarge |

#### Component Spacing
| Token | Value | Usage |
|-------|-------|-------|
| `--space-page` | 32px | Page padding |
| `--space-section` | 24px | Section gaps |
| `--space-gap` | 20px | Card gaps |
| `--space-component` | 16px | Component internal |

---

### 2.4 Border Radius

| Token | Value | Usage |
|-------|-------|-------|
| `--radius-xs` | 4px | Small elements |
| `--radius-sm` | 6px | Buttons small |
| `--radius-md` | 8px | Inputs |
| `--radius-lg` | 12px | Cards |
| `--radius-xl` | 16px | Modals |
| `--radius-2xl` | 20px | Large cards |
| `--radius-3xl` | 24px | Hero cards |
| `--radius-full` | 9999px | Pills, avatars |

---

### 2.5 Shadows

#### Light Mode
| Token | Value |
|-------|-------|
| `--shadow-xs` | 0 1px 2px 0 rgb(0 0 0 / 0.04) |
| `--shadow-sm` | 0 1px 3px 0 rgb(0 0 0 / 0.06), 0 1px 2px -1px rgb(0 0 0 / 0.04) |
| `--shadow-md` | 0 4px 6px -1px rgb(0 0 0 / 0.06), 0 2px 4px -2px rgb(0 0 0 / 0.04) |
| `--shadow-lg` | 0 10px 15px -3px rgb(0 0 0 / 0.08), 0 4px 6px -4px rgb(0 0 0 / 0.04) |
| `--shadow-xl` | 0 20px 25px -5px rgb(0 0 0 / 0.08), 0 8px 10px -6px rgb(0 0 0 / 0.04) |
| `--shadow-2xl` | 0 25px 50px -12px rgb(0 0 0 / 0.15) |

#### Dark Mode
Shadows use higher opacity (0.2-0.4) for better visibility on dark backgrounds.

---

### 2.6 Transitions

| Token | Duration | Usage |
|-------|---------|-------|
| `--transition-instant` | 50ms | Immediate feedback |
| `--transition-fast` | 150ms | Hover states |
| `--transition-base` | 200ms | Default transitions |
| `--transition-slow` | 300ms | Modal/overlay |
| `--transition-slower` | 400ms | Large animations |

---

### 2.7 Z-Index Scale

| Token | Value | Usage |
|-------|-------|-------|
| `--z-dropdown` | 50 | Dropdown menus |
| `--z-sticky` | 100 | Sticky headers |
| `--z-fixed` | 200 | Fixed elements |
| `--z-modal-backdrop` | 300 | Modal overlays |
| `--z-modal` | 400 | Modal dialogs |
| `--z-popover` | 500 | Popovers |
| `--z-tooltip` | 600 | Tooltips |
| `--z-toast` | 700 | Toast notifications |

---

## 3. Layout System

### 3.1 Page Layout
- **Sidebar Width**: 260px (expanded), 80px (collapsed)
- **Page Padding**: 32px
- **Section Gap**: 24px
- **Content Max Width**: 1400px

### 3.2 Grid System
- **Stat Cards**: 3 columns on desktop, 2 on tablet, 1 on mobile
- **Dashboard Grid**: Flexible with minimum 280px card width
- **Article Grid**: Auto-fill with minmax(280px, 1fr)

### 3.3 Responsive Breakpoints
| Breakpoint | Width | Layout |
|------------|-------|--------|
| Desktop XL | 1400px+ | Full sidebar, 3-column stats |
| Desktop | 1200px+ | Collapsed sidebar |
| Tablet | 1024px+ | Adjusted grids |
| Mobile | 768px+ | Stacked layout |
| Mobile S | 640px+ | Single column |

---

## 4. Component Guidelines

### 4.1 Buttons

#### Primary Button
```css
background: var(--color-primary);
color: #ffffff;
border-radius: var(--radius-lg);
padding: 12px 20px;
font-weight: var(--font-weight-semibold);
box-shadow: var(--shadow-sm);
```
- Hover: `--color-primary-hover`, shadow increase
- Active: translateY(0)
- Disabled: opacity 0.6

#### Secondary Button
```css
background: var(--color-surface);
color: var(--color-text-primary);
border: 1px solid var(--color-border);
```
- Hover: background `--color-surface-alt`

### 4.2 Cards

#### Base Card
```css
background: var(--color-surface);
border-radius: var(--radius-xl);
border: 1px solid var(--color-border);
box-shadow: var(--shadow-sm);
padding: var(--space-6);
```
- Hover: box-shadow to `--shadow-lg`, translateY(-2px)

### 4.3 Form Inputs

```css
background: var(--color-surface-alt);
border: 1px solid var(--color-border);
border-radius: var(--radius-lg);
padding: 12px 16px;
font-size: var(--font-size-base);
```
- Focus: border-color `--color-primary`, box-shadow 0 0 0 3px `--color-primary-soft`

### 4.4 Sidebar Navigation

- Fixed position with 16px margin from edges
- 260px width (80px collapsed)
- Border radius: `--radius-xl`
- Active item: `--color-primary-soft` background

---

## 5. Accessibility

### 5.1 Color Contrast
- Primary text: minimum 4.5:1 contrast ratio
- Large text (18px+): minimum 3:1 contrast ratio
- UI components: minimum 3:1 contrast ratio

### 5.2 Focus States
```css
input:focus-visible,
button:focus-visible {
  outline: 2px solid var(--color-primary);
  outline-offset: 2px;
}
```

### 5.3 Motion Preferences
```css
@media (prefers-reduced-motion: reduce) {
  * {
    animation-duration: 0.01ms !important;
    transition-duration: 0.01ms !important;
  }
}
```

### 5.4 Screen Reader Support
- Semantic HTML elements
- ARIA labels where needed
- `.sr-only` class for visually hidden text

---

## 6. Icon System

### 6.1 Icon Library
Using inline SVG icons with consistent:
- **Size**: 20x20px viewBox
- **Stroke Width**: 2px
- **Style**: Lucide icons (rounded caps and joins)

### 6.2 Icon Usage
- Navigation: 20px
- Buttons: 16-18px
- Cards: 20px
- Badges: 12-14px

---

## 7. Animation Guidelines

### 7.1 Micro-interactions
- **Hover lift**: `translateY(-2px)` with shadow increase
- **Button press**: `translateY(0)` on active
- **Card hover**: Subtle lift with shadow enhancement

### 7.2 Transitions
- Duration: 150-200ms for UI feedback
- Easing: `ease` for most transitions
- Avoid: Motion that causes layout shifts

### 7.3 Loading States
- Skeleton shimmer animation: 1.2s infinite
- Spinner animation: 1.2s linear infinite
- Pulse animation: 1.5s infinite for indicators

### 7.4 Vue Transition System

#### Page Transitions
```css
.page-enter-active,
.page-leave-active {
  transition: opacity 0.3s ease, transform 0.3s ease;
}

.page-enter-from {
  opacity: 0;
  transform: translateY(12px);
}

.page-leave-to {
  opacity: 0;
  transform: translateY(-8px);
}
```

#### List Animations (CRUD)
```html
<TransitionGroup name="list" tag="div" class="grid">
  <article v-for="item in items" :key="item.id">
    <!-- item content -->
  </article>
</TransitionGroup>
```

```css
.list-enter-active,
.list-leave-active {
  transition: all 0.3s ease;
}

.list-enter-from {
  opacity: 0;
  transform: translateX(-24px);
}

.list-leave-to {
  opacity: 0;
  transform: translateX(24px);
}

.list-leave-active {
  position: absolute;
}

.list-move {
  transition: transform 0.3s ease;
}
```

#### Toast Notifications
```css
.toast-enter-active,
.toast-leave-active {
  transition: opacity 0.3s ease, transform 0.3s ease;
}

.toast-enter-from {
  opacity: 0;
  transform: translateY(24px) scale(0.95);
}

.toast-leave-to {
  opacity: 0;
  transform: translateY(-12px) scale(0.95);
}
```

#### Utility Animation Classes
| Class | Description |
|-------|-------------|
| `.fade-enter-active` | Fade in/out transitions |
| `.slide-up-enter-active` | Slide up animations |
| `.slide-down-enter-active` | Slide down animations |
| `.slide-left-enter-active` | Slide left animations |
| `.slide-right-enter-active` | Slide right animations |
| `.scale-enter-active` | Scale up/down transitions |
| `.animate-pulse-soft` | Soft pulse animation |
| `.animate-spin-smooth` | Smooth spin animation |
| `.animate-bounce-in` | Bounce in animation |

---

## 8. File Structure

```
src/
├── styles/
│   └── variables.css     # Design tokens and global styles
├── components/
│   └── admin/
│       ├── AdminSidebar.vue
│       ├── StatCard.vue
│       ├── TodoCard.vue
│       ├── TrendChart.vue
│       └── HeatmapChart.vue
└── views/
    └── admin/
        ├── DashboardView.vue
        ├── LoginView.vue
        ├── ArticleListView.vue
        └── ArticleEditorView.vue
```

---

## 9. Implementation Checklist

### Visual Quality
- [x] No emojis used as icons (SVG icons only)
- [x] Consistent icon sizing (20x20 viewBox)
- [x] Hover states with smooth transitions (150-200ms)
- [x] Brand colors verified (NectarPin blue gradient)

### Interaction
- [x] `cursor-pointer` on all clickable elements
- [x] Clear hover feedback on interactive elements
- [x] Focus states visible for keyboard navigation
- [x] Active states for buttons

### Light/Dark Mode
- [x] Text contrast meets WCAG AA (4.5:1 minimum)
- [x] Surface colors work in both modes
- [x] Border visibility in both modes
- [x] Uses `prefers-color-scheme` media query

### Layout
- [x] Responsive at 375px, 768px, 1024px, 1440px
- [x] Proper spacing with CSS variables
- [x] No content hidden behind fixed elements
- [x] Floating navbar with proper margin

### Accessibility
- [x] Focus states with outline
- [x] Reduced motion support
- [x] Semantic HTML structure
- [x] ARIA labels on icon-only buttons

---

## 10. Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2026-03-19 | Initial design system with complete token set |
