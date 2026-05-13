# Terminoter

A cute, simple note-taking app for the terminal.

>[!NOTE]
> **AI Disclosure:**
> None of the contents of this repo was **AI** generated (not that there's anything wrong with that).


Demo:
![simple_demo](./demos/banner.webp)


## Roadmap

- [x] Add/remove notes.
- [x] Support all kind of text (emoji, lists, etc).
- [x] Save to default location (usually `~/.local/share/terminoter/...`).
  - [x] Optionally read from/write to custom file.
- [ ] Different layouts according to terminal size.
  - [ ] Handle the case where the notes occupy more than the available space.
- [ ] Change layout.
- [ ] Collapse notes, showing only first line or title.
- [ ] Delete everything.
- [ ] Atomic writes.
- [ ] Edit notes.
  - [ ] Minimal edition facilities (i.e. emacs/vim binidngs).
    - [ ] Delete word (forwards/backwards).
    - [ ] Delete line (id).
    - [ ] Delete list, paragraph, etc.
    - [ ] Switch case.
    - [ ] ...
  - [ ] Minimal `terminal` markup (i.e. add `bold, italic, underline`, etc).
- [ ] Change notes' order.
- [ ] Save metadata (i.e. date/time).
- [ ] Backup file.
- [ ] Emoji chooser.
  - [ ] Icon/symbol chooser (Unicode/Nerdfont/other icons).
- [ ] Rewrite it in Rust :crab:.
