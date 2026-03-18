; ModuleID = 'main'
source_filename = "main"

define i32 @foo() {
entry:
  ret i32 0
}

define i32 @main() {
entry:
  %x = alloca float, align 4
  store i32 7, ptr %x, align 4
  store i32 1, ptr %x, align 4
  ret i32 0
}
