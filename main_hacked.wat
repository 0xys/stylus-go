(module
  (type (;0;) (func (param i32 i32 i32 i32) (result i32)))
  (type (;1;) (func (param i32)))
  (type (;2;) (func (param i32 i32) (result i32)))
  (type (;3;) (func (param i32) (result i32)))
  (type (;4;) (func (param i32 i32)))
  (type (;5;) (func))
  (type (;6;) (func (result i64)))
  (type (;7;) (func (param i32 i32 i32) (result i32)))
  (type (;8;) (func (result i32)))
  (type (;9;) (func (param i32 i32 i32 i32)))
  (import "vm_hooks" "memory_grow" (func $memory_grow (type 1)))
  (import "vm_hooks" "chainid" (func $chainid (type 6)))
  (func $memset (type 2) (param i32 i32) (result i32)
    (local i32 i32)
    local.get 1
    i32.const 33
    i32.ge_u
    if  ;; label = @1
      local.get 0
      i32.const 0
      local.get 1
      memory.fill
      local.get 0
      return
    end
    block  ;; label = @1
      local.get 1
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      i32.const 0
      i32.store8
      local.get 0
      local.get 1
      i32.add
      local.tee 2
      i32.const 1
      i32.sub
      i32.const 0
      i32.store8
      local.get 1
      i32.const 3
      i32.lt_u
      br_if 0 (;@1;)
      local.get 0
      i32.const 0
      i32.store8 offset=2
      local.get 0
      i32.const 0
      i32.store8 offset=1
      local.get 2
      i32.const 3
      i32.sub
      i32.const 0
      i32.store8
      local.get 2
      i32.const 2
      i32.sub
      i32.const 0
      i32.store8
      local.get 1
      i32.const 7
      i32.lt_u
      br_if 0 (;@1;)
      local.get 0
      i32.const 0
      i32.store8 offset=3
      local.get 2
      i32.const 4
      i32.sub
      i32.const 0
      i32.store8
      local.get 1
      i32.const 9
      i32.lt_u
      br_if 0 (;@1;)
      local.get 0
      i32.const 0
      local.get 0
      i32.sub
      i32.const 3
      i32.and
      local.tee 3
      i32.add
      local.tee 2
      i32.const 0
      i32.store
      local.get 2
      local.get 1
      local.get 3
      i32.sub
      i32.const -4
      i32.and
      local.tee 3
      i32.add
      local.tee 1
      i32.const 4
      i32.sub
      i32.const 0
      i32.store
      local.get 3
      i32.const 9
      i32.lt_u
      br_if 0 (;@1;)
      local.get 2
      i32.const 0
      i32.store offset=8
      local.get 2
      i32.const 0
      i32.store offset=4
      local.get 1
      i32.const 8
      i32.sub
      i32.const 0
      i32.store
      local.get 1
      i32.const 12
      i32.sub
      i32.const 0
      i32.store
      local.get 3
      i32.const 25
      i32.lt_u
      br_if 0 (;@1;)
      local.get 2
      i32.const 0
      i32.store offset=24
      local.get 2
      i32.const 0
      i32.store offset=20
      local.get 2
      i32.const 0
      i32.store offset=16
      local.get 2
      i32.const 0
      i32.store offset=12
      local.get 1
      i32.const 16
      i32.sub
      i32.const 0
      i32.store
      local.get 1
      i32.const 20
      i32.sub
      i32.const 0
      i32.store
      local.get 1
      i32.const 24
      i32.sub
      i32.const 0
      i32.store
      local.get 1
      i32.const 28
      i32.sub
      i32.const 0
      i32.store
      local.get 3
      local.get 2
      i32.const 4
      i32.and
      i32.const 24
      i32.or
      local.tee 1
      i32.sub
      local.tee 3
      i32.const 32
      i32.lt_u
      br_if 0 (;@1;)
      local.get 1
      local.get 2
      i32.add
      local.set 1
      loop  ;; label = @2
        local.get 1
        i64.const 0
        i64.store offset=24
        local.get 1
        i64.const 0
        i64.store offset=16
        local.get 1
        i64.const 0
        i64.store offset=8
        local.get 1
        i64.const 0
        i64.store
        local.get 1
        i32.const 32
        i32.add
        local.set 1
        local.get 3
        i32.const 32
        i32.sub
        local.tee 3
        i32.const 31
        i32.gt_u
        br_if 0 (;@2;)
      end
    end
    local.get 0)
  (func $memcpy (type 7) (param i32 i32 i32) (result i32)
    (local i32 i32 i32 i32)
    block  ;; label = @1
      block (result i32)  ;; label = @2
        block  ;; label = @3
          local.get 2
          i32.const 32
          i32.le_u
          if  ;; label = @4
            local.get 1
            i32.const 3
            i32.and
            i32.eqz
            local.get 2
            i32.eqz
            i32.or
            br_if 1 (;@3;)
            local.get 0
            local.get 1
            i32.load8_u
            i32.store8
            local.get 0
            i32.const 1
            i32.add
            local.get 1
            i32.const 1
            i32.add
            local.tee 3
            i32.const 3
            i32.and
            i32.eqz
            local.get 2
            i32.const 1
            i32.sub
            local.tee 5
            i32.eqz
            i32.or
            br_if 2 (;@2;)
            drop
            local.get 0
            local.get 1
            i32.load8_u offset=1
            i32.store8 offset=1
            local.get 0
            i32.const 2
            i32.add
            local.get 1
            i32.const 2
            i32.add
            local.tee 3
            i32.const 3
            i32.and
            i32.eqz
            local.get 2
            i32.const 2
            i32.sub
            local.tee 5
            i32.eqz
            i32.or
            br_if 2 (;@2;)
            drop
            local.get 0
            local.get 1
            i32.load8_u offset=2
            i32.store8 offset=2
            local.get 0
            i32.const 3
            i32.add
            local.get 1
            i32.const 3
            i32.add
            local.tee 3
            i32.const 3
            i32.and
            i32.eqz
            local.get 2
            i32.const 3
            i32.sub
            local.tee 5
            i32.eqz
            i32.or
            br_if 2 (;@2;)
            drop
            local.get 0
            local.get 1
            i32.load8_u offset=3
            i32.store8 offset=3
            local.get 2
            i32.const 4
            i32.sub
            local.set 5
            local.get 1
            i32.const 4
            i32.add
            local.set 3
            local.get 0
            i32.const 4
            i32.add
            br 2 (;@2;)
          end
          local.get 0
          local.get 1
          local.get 2
          memory.copy
          local.get 0
          return
        end
        local.get 2
        local.set 5
        local.get 1
        local.set 3
        local.get 0
      end
      local.tee 4
      i32.const 3
      i32.and
      local.tee 1
      i32.eqz
      if  ;; label = @2
        block  ;; label = @3
          local.get 5
          i32.const 16
          i32.lt_u
          if  ;; label = @4
            local.get 5
            local.set 2
            br 1 (;@3;)
          end
          local.get 5
          i32.const 16
          i32.sub
          local.tee 2
          i32.const 16
          i32.and
          i32.eqz
          if  ;; label = @4
            local.get 4
            local.get 3
            i64.load align=4
            i64.store align=4
            local.get 4
            local.get 3
            i64.load offset=8 align=4
            i64.store offset=8 align=4
            local.get 4
            i32.const 16
            i32.add
            local.set 4
            local.get 3
            i32.const 16
            i32.add
            local.set 3
            local.get 2
            local.set 5
          end
          local.get 2
          i32.const 16
          i32.lt_u
          br_if 0 (;@3;)
          local.get 5
          local.set 2
          loop  ;; label = @4
            local.get 4
            local.get 3
            i64.load align=4
            i64.store align=4
            local.get 4
            local.get 3
            i64.load offset=8 align=4
            i64.store offset=8 align=4
            local.get 4
            local.get 3
            i64.load offset=16 align=4
            i64.store offset=16 align=4
            local.get 4
            local.get 3
            i64.load offset=24 align=4
            i64.store offset=24 align=4
            local.get 4
            i32.const 32
            i32.add
            local.set 4
            local.get 3
            i32.const 32
            i32.add
            local.set 3
            local.get 2
            i32.const 32
            i32.sub
            local.tee 2
            i32.const 15
            i32.gt_u
            br_if 0 (;@4;)
          end
        end
        local.get 2
        i32.const 8
        i32.ge_u
        if  ;; label = @3
          local.get 4
          local.get 3
          i64.load align=4
          i64.store align=4
          local.get 4
          i32.const 8
          i32.add
          local.set 4
          local.get 3
          i32.const 8
          i32.add
          local.set 3
        end
        local.get 2
        i32.const 4
        i32.and
        if  ;; label = @3
          local.get 4
          local.get 3
          i32.load
          i32.store
          local.get 4
          i32.const 4
          i32.add
          local.set 4
          local.get 3
          i32.const 4
          i32.add
          local.set 3
        end
        local.get 2
        i32.const 2
        i32.and
        if  ;; label = @3
          local.get 4
          local.get 3
          i32.load16_u align=1
          i32.store16 align=1
          local.get 4
          i32.const 2
          i32.add
          local.set 4
          local.get 3
          i32.const 2
          i32.add
          local.set 3
        end
        local.get 2
        i32.const 1
        i32.and
        i32.eqz
        br_if 1 (;@1;)
        local.get 4
        local.get 3
        i32.load8_u
        i32.store8
        local.get 0
        return
      end
      block  ;; label = @2
        block  ;; label = @3
          block (result i32)  ;; label = @4
            block  ;; label = @5
              local.get 5
              i32.const 32
              i32.ge_u
              if  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    local.get 1
                    i32.const 1
                    i32.sub
                    br_table 3 (;@5;) 0 (;@8;) 1 (;@7;) 7 (;@1;)
                  end
                  local.get 4
                  local.get 3
                  i32.load
                  i32.store16 align=1
                  local.get 4
                  local.get 3
                  i32.const 2
                  i32.add
                  i32.load align=2
                  i32.store offset=2
                  local.get 4
                  local.get 3
                  i32.const 6
                  i32.add
                  i64.load align=2
                  i64.store offset=6 align=4
                  local.get 3
                  i32.const 18
                  i32.add
                  local.set 1
                  i32.const 14
                  local.set 6
                  local.get 3
                  i32.const 14
                  i32.add
                  i32.load align=2
                  local.set 3
                  i32.const 14
                  local.set 5
                  local.get 4
                  i32.const 18
                  i32.add
                  br 3 (;@4;)
                end
                local.get 4
                local.get 3
                i32.load
                i32.store8
                local.get 4
                local.get 3
                i32.const 1
                i32.add
                i32.load align=1
                i32.store offset=1
                local.get 4
                local.get 3
                i32.const 5
                i32.add
                i64.load align=1
                i64.store offset=5 align=4
                local.get 3
                i32.const 17
                i32.add
                local.set 1
                i32.const 13
                local.set 6
                local.get 3
                i32.const 13
                i32.add
                i32.load align=1
                local.set 3
                i32.const 15
                local.set 5
                local.get 4
                i32.const 17
                i32.add
                br 2 (;@4;)
              end
              block (result i32)  ;; label = @6
                local.get 5
                i32.const 16
                i32.lt_u
                if  ;; label = @7
                  local.get 4
                  local.set 2
                  local.get 3
                  br 1 (;@6;)
                end
                local.get 4
                local.get 3
                i32.load8_u
                i32.store8
                local.get 4
                local.get 3
                i32.load offset=1 align=1
                i32.store offset=1 align=1
                local.get 4
                local.get 3
                i64.load offset=5 align=1
                i64.store offset=5 align=1
                local.get 4
                local.get 3
                i32.load16_u offset=13 align=1
                i32.store16 offset=13 align=1
                local.get 4
                local.get 3
                i32.load8_u offset=15
                i32.store8 offset=15
                local.get 4
                i32.const 16
                i32.add
                local.set 2
                local.get 3
                i32.const 16
                i32.add
              end
              local.set 1
              local.get 5
              i32.const 8
              i32.and
              br_if 2 (;@3;)
              br 3 (;@2;)
            end
            local.get 4
            local.get 3
            i32.load
            local.tee 1
            i32.store8
            local.get 4
            local.get 1
            i32.const 16
            i32.shr_u
            i32.store8 offset=2
            local.get 4
            local.get 1
            i32.const 8
            i32.shr_u
            i32.store8 offset=1
            local.get 4
            local.get 3
            i32.const 3
            i32.add
            i32.load align=1
            i32.store offset=3
            local.get 4
            local.get 3
            i32.const 7
            i32.add
            i64.load align=1
            i64.store offset=7 align=4
            local.get 3
            i32.const 19
            i32.add
            local.set 1
            i32.const 15
            local.set 6
            local.get 3
            i32.const 15
            i32.add
            i32.load align=1
            local.set 3
            i32.const 13
            local.set 5
            local.get 4
            i32.const 19
            i32.add
          end
          local.set 2
          local.get 4
          local.get 6
          i32.add
          local.get 3
          i32.store
        end
        local.get 2
        local.get 1
        i64.load align=1
        i64.store align=1
        local.get 2
        i32.const 8
        i32.add
        local.set 2
        local.get 1
        i32.const 8
        i32.add
        local.set 1
      end
      local.get 5
      i32.const 4
      i32.and
      if  ;; label = @2
        local.get 2
        local.get 1
        i32.load align=1
        i32.store align=1
        local.get 2
        i32.const 4
        i32.add
        local.set 2
        local.get 1
        i32.const 4
        i32.add
        local.set 1
      end
      local.get 5
      i32.const 2
      i32.and
      if  ;; label = @2
        local.get 2
        local.get 1
        i32.load16_u align=1
        i32.store16 align=1
        local.get 2
        i32.const 2
        i32.add
        local.set 2
        local.get 1
        i32.const 2
        i32.add
        local.set 1
      end
      local.get 5
      i32.const 1
      i32.and
      i32.eqz
      br_if 0 (;@1;)
      local.get 2
      local.get 1
      i32.load8_u
      i32.store8
    end
    local.get 0)
  (func $runtime.memequal (type 0) (param i32 i32 i32 i32) (result i32)
    (local i32 i32)
    i32.const 0
    local.set 3
    block (result i32)  ;; label = @1
      loop  ;; label = @2
        local.get 2
        local.get 2
        local.get 3
        i32.eq
        br_if 1 (;@1;)
        drop
        local.get 1
        local.get 3
        i32.add
        local.set 4
        local.get 0
        local.get 3
        i32.add
        local.get 3
        i32.const 1
        i32.add
        local.set 3
        i32.load8_u
        local.get 4
        i32.load8_u
        i32.eq
        br_if 0 (;@2;)
      end
      local.get 3
      i32.const 1
      i32.sub
    end
    local.get 2
    i32.ge_u)
  (func $runtime.hash32 (type 0) (param i32 i32 i32 i32) (result i32)
    local.get 2
    i32.const -2128831035
    i32.mul
    local.set 2
    loop  ;; label = @1
      local.get 1
      if  ;; label = @2
        local.get 1
        i32.const 1
        i32.sub
        local.set 1
        local.get 2
        local.get 0
        i32.load8_u
        i32.xor
        i32.const 16777619
        i32.mul
        local.set 2
        local.get 0
        i32.const 1
        i32.add
        local.set 0
        br 1 (;@1;)
      end
    end
    local.get 2)
  (func $malloc (type 3) (param i32) (result i32)
    (local i32 i32)
    global.get $__stack_pointer
    i32.const 16
    i32.sub
    local.tee 1
    global.set $__stack_pointer
    block  ;; label = @1
      local.get 0
      i32.eqz
      if  ;; label = @2
        br 1 (;@1;)
      end
      local.get 0
      i32.const 0
      i32.ge_s
      if  ;; label = @2
        local.get 0
        call $runtime.alloc
        local.set 2
        local.get 1
        local.get 0
        i32.store offset=8
        local.get 1
        local.get 0
        i32.store offset=4
        local.get 1
        local.get 2
        i32.store
        local.get 1
        local.get 2
        i32.store offset=12
        local.get 1
        i32.const 12
        i32.add
        local.get 1
        call $runtime.hashmapBinarySet
        br 1 (;@1;)
      end
      call $runtime.slicePanic
      unreachable
    end
    local.get 1
    i32.const 16
    i32.add
    global.set $__stack_pointer
    local.get 2)
  (func $runtime.alloc (type 3) (param i32) (result i32)
    (local i32 i32 i32)
    i32.const 65896
    i32.const 65896
    i64.load
    i64.const 1
    i64.add
    i64.store
    i32.const 65708
    i32.const 65708
    i32.load
    local.tee 3
    local.get 0
    i32.const 15
    i32.add
    i32.const -16
    i32.and
    local.tee 1
    i32.add
    local.tee 0
    i32.store
    i32.const 65888
    i32.const 65888
    i64.load
    local.get 1
    i64.extend_i32_u
    i64.add
    i64.store
    i32.const 65760
    i32.load
    local.set 2
    block  ;; label = @1
      loop  ;; label = @2
        local.get 0
        local.get 2
        i32.lt_u
        br_if 1 (;@1;)
        memory.size
        memory.grow
        i32.const -1
        i32.ne
        if  ;; label = @3
          i32.const 65760
          memory.size
          i32.const 16
          i32.shl
          local.tee 2
          i32.store
          i32.const 65708
          i32.load
          local.set 0
          br 1 (;@2;)
        end
      end
      i32.const 65600
      i32.const 13
      call $runtime.runtimePanicAt
      unreachable
    end
    local.get 3
    local.get 1
    call $memset)
  (func $runtime.hashmapBinarySet (type 4) (param i32 i32)
    i32.const 65720
    local.get 0
    local.get 1
    local.get 0
    i32.const 65732
    i32.load
    i32.const 65724
    i32.load
    local.get 0
    call $runtime.hash32
    call $runtime.hashmapSet)
  (func $runtime.slicePanic (type 5)
    i32.const 65683
    i32.const 18
    call $runtime.runtimePanicAt
    unreachable)
  (func $free (type 1) (param i32)
    (local i32)
    global.get $__stack_pointer
    i32.const 16
    i32.sub
    local.tee 1
    global.set $__stack_pointer
    block  ;; label = @1
      local.get 0
      if  ;; label = @2
        local.get 1
        local.get 0
        i32.store offset=12
        local.get 1
        i32.const 12
        i32.add
        local.get 1
        call $runtime.hashmapBinaryGet
        i32.const 1
        i32.and
        i32.eqz
        br_if 1 (;@1;)
        local.get 1
        local.get 0
        i32.store
        local.get 1
        call $runtime.hashmapBinaryDelete
      end
      local.get 1
      i32.const 16
      i32.add
      global.set $__stack_pointer
      return
    end
    i32.const 65560
    call $runtime._panic
    unreachable)
  (func $runtime.hashmapBinaryGet (type 2) (param i32 i32) (result i32)
    i32.const 65720
    local.get 0
    local.get 1
    local.get 0
    i32.const 65732
    i32.load
    i32.const 65724
    i32.load
    local.get 0
    call $runtime.hash32
    call $runtime.hashmapGet)
  (func $runtime.hashmapBinaryDelete (type 1) (param i32)
    (local i32 i32 i32 i32 i32 i32 i32)
    local.get 0
    i32.const 65732
    i32.load
    i32.const 65724
    i32.load
    i32.const 0
    call $runtime.hash32
    local.tee 1
    i32.const 24
    i32.shr_u
    local.tee 2
    i32.const 1
    local.get 2
    select
    local.set 3
    i32.const 65720
    local.get 1
    call $runtime.hashmapBucketAddrForHash
    local.set 1
    block  ;; label = @1
      loop  ;; label = @2
        local.get 1
        i32.eqz
        br_if 1 (;@1;)
        i32.const 0
        local.set 2
        block  ;; label = @3
          loop  ;; label = @4
            local.get 2
            i32.const 8
            i32.ne
            if  ;; label = @5
              block  ;; label = @6
                local.get 1
                local.get 2
                i32.add
                local.tee 4
                i32.load8_u
                local.get 3
                i32.ne
                br_if 0 (;@6;)
                i32.const 65748
                i32.load
                local.tee 5
                i32.eqz
                br_if 3 (;@3;)
                local.get 0
                i32.const 65732
                i32.load
                local.tee 6
                local.get 2
                i32.mul
                local.get 1
                i32.add
                i32.const 12
                i32.add
                local.tee 7
                local.get 6
                i32.const 65744
                i32.load
                local.get 5
                call_indirect (type 0)
                i32.const 1
                i32.and
                i32.eqz
                br_if 0 (;@6;)
                local.get 4
                i32.const 0
                i32.store8
                local.get 7
                i32.const 65732
                i32.load
                call $memset
                drop
                i32.const 65736
                i32.load
                local.tee 0
                local.get 2
                i32.mul
                i32.const 65732
                i32.load
                i32.const 3
                i32.shl
                i32.add
                local.get 1
                i32.add
                i32.const 12
                i32.add
                local.get 0
                call $memset
                drop
                i32.const 65728
                i32.const 65728
                i32.load
                i32.const 1
                i32.sub
                i32.store
                br 5 (;@1;)
              end
              local.get 2
              i32.const 1
              i32.add
              local.set 2
              br 1 (;@4;)
            end
          end
          local.get 1
          i32.load offset=8
          local.set 1
          br 1 (;@2;)
        end
      end
      call $runtime.nilPanic
      unreachable
    end)
  (func $runtime._panic (type 1) (param i32)
    i32.const 65613
    i32.const 7
    call $runtime.printstring
    local.get 0
    i32.load
    local.get 0
    i32.load offset=4
    call $runtime.printstring
    i32.const 10
    call $runtime.putchar
    unreachable)
  (func $user_entrypoint (type 8) (result i32)
    i32.const 0
    call $memory_grow
    call $chainid
    drop
    i32.const 0)
  (func $runtime.hashmapGet (type 0) (param i32 i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32)
    local.get 3
    i32.const 24
    i32.shr_u
    local.tee 4
    i32.const 1
    local.get 4
    select
    local.set 6
    local.get 0
    local.get 3
    call $runtime.hashmapBucketAddrForHash
    local.set 4
    block  ;; label = @1
      block  ;; label = @2
        loop  ;; label = @3
          local.get 4
          if  ;; label = @4
            i32.const 0
            local.set 3
            loop  ;; label = @5
              local.get 3
              i32.const 8
              i32.ne
              if  ;; label = @6
                block  ;; label = @7
                  local.get 3
                  local.get 4
                  i32.add
                  i32.load8_u
                  local.get 6
                  i32.ne
                  br_if 0 (;@7;)
                  local.get 0
                  i32.load offset=28
                  local.tee 7
                  i32.eqz
                  br_if 6 (;@1;)
                  local.get 0
                  i32.load offset=16
                  local.set 8
                  local.get 1
                  local.get 0
                  i32.load offset=12
                  local.tee 5
                  local.get 3
                  i32.mul
                  local.get 4
                  i32.add
                  i32.const 12
                  i32.add
                  local.get 5
                  local.get 0
                  i32.load offset=24
                  local.get 7
                  call_indirect (type 0)
                  i32.const 1
                  i32.and
                  i32.eqz
                  br_if 0 (;@7;)
                  local.get 2
                  local.get 3
                  local.get 8
                  i32.mul
                  local.get 5
                  i32.const 3
                  i32.shl
                  i32.add
                  local.get 4
                  i32.add
                  i32.const 12
                  i32.add
                  local.get 0
                  i32.load offset=16
                  call $memcpy
                  drop
                  br 5 (;@2;)
                end
                local.get 3
                i32.const 1
                i32.add
                local.set 3
                br 1 (;@5;)
              end
            end
            local.get 4
            i32.load offset=8
            local.set 4
            br 1 (;@3;)
          end
        end
        local.get 2
        local.get 0
        i32.load offset=16
        call $memset
        drop
      end
      local.get 4
      i32.const 0
      i32.ne
      return
    end
    call $runtime.nilPanic
    unreachable)
  (func $runtime.hashmapBucketAddrForHash (type 2) (param i32 i32) (result i32)
    local.get 0
    i32.load
    local.get 0
    i32.load offset=16
    local.get 0
    i32.load offset=12
    i32.add
    i32.const 3
    i32.shl
    i32.const 12
    i32.add
    i32.const -1
    i32.const -1
    local.get 0
    i32.load8_u offset=20
    local.tee 0
    i32.shl
    i32.const -1
    i32.xor
    local.get 0
    i32.const 31
    i32.gt_u
    select
    local.get 1
    i32.and
    i32.mul
    i32.add)
  (func $runtime.nilPanic (type 5)
    i32.const 65642
    i32.const 23
    call $runtime.runtimePanicAt
    unreachable)
  (func $runtime.printstring (type 4) (param i32 i32)
    local.get 1
    i32.const 0
    local.get 1
    i32.const 0
    i32.gt_s
    select
    local.set 1
    loop  ;; label = @1
      local.get 1
      if  ;; label = @2
        local.get 0
        i32.load8_u
        call $runtime.putchar
        local.get 1
        i32.const 1
        i32.sub
        local.set 1
        local.get 0
        i32.const 1
        i32.add
        local.set 0
        br 1 (;@1;)
      end
    end)
  (func $runtime.hashmapSet (type 9) (param i32 i32 i32 i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get $__stack_pointer
    i32.const 48
    i32.sub
    local.tee 4
    global.set $__stack_pointer
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        i32.eqz
        br_if 0 (;@2;)
        block  ;; label = @3
          local.get 0
          i32.load8_u offset=20
          local.tee 5
          i32.const 29
          i32.gt_u
          br_if 0 (;@3;)
          local.get 0
          i32.load offset=8
          i32.const 6
          local.get 5
          i32.shl
          i32.le_u
          br_if 0 (;@3;)
          i32.const 65704
          i32.const 65704
          i32.load
          local.tee 3
          i32.const 7
          i32.shl
          local.get 3
          i32.xor
          local.tee 3
          i32.const 1
          i32.shr_u
          local.get 3
          i32.xor
          local.tee 3
          i32.const 9
          i32.shl
          local.get 3
          i32.xor
          local.tee 3
          i32.store
          local.get 4
          i64.const 0
          i64.store offset=24
          local.get 4
          i32.const 0
          i32.store offset=16
          local.get 4
          local.get 0
          i64.load offset=32 align=4
          i64.store offset=40
          local.get 4
          local.get 0
          i64.load offset=24 align=4
          i64.store offset=32
          local.get 4
          local.get 0
          i32.load offset=16
          local.tee 6
          i32.store offset=24
          local.get 4
          local.get 0
          i32.load offset=12
          local.tee 7
          i32.store offset=20
          local.get 4
          local.get 3
          i32.store offset=12
          local.get 4
          local.get 5
          i32.const 1
          i32.add
          local.tee 3
          i32.store8 offset=28
          local.get 4
          local.get 6
          local.get 7
          i32.add
          i32.const 3
          i32.shl
          i32.const 12
          i32.add
          local.get 3
          i32.shl
          call $runtime.alloc
          i32.store offset=8
          local.get 0
          i32.load offset=12
          call $runtime.alloc
          local.set 13
          local.get 0
          i32.load offset=16
          call $runtime.alloc
          local.set 11
          i32.const 0
          local.set 5
          i32.const 0
          local.set 3
          i32.const 0
          local.set 6
          i32.const 0
          local.set 7
          loop  ;; label = @4
            local.get 7
            i32.eqz
            if  ;; label = @5
              i32.const 0
              i32.const 1
              local.get 0
              i32.load8_u offset=20
              local.tee 7
              i32.shl
              local.get 7
              i32.const 31
              i32.gt_u
              select
              local.set 12
              local.get 0
              i32.load
              local.set 7
            end
            block  ;; label = @5
              loop  ;; label = @6
                local.get 5
                i32.const 255
                i32.and
                i32.const 8
                i32.ge_u
                if  ;; label = @7
                  local.get 3
                  i32.eqz
                  br_if 5 (;@2;)
                  i32.const 0
                  local.set 5
                  local.get 3
                  i32.load offset=8
                  local.set 3
                end
                local.get 3
                i32.eqz
                if  ;; label = @7
                  local.get 6
                  local.get 12
                  i32.ge_u
                  br_if 2 (;@5;)
                  local.get 7
                  local.get 0
                  i32.load offset=16
                  local.get 0
                  i32.load offset=12
                  i32.add
                  i32.const 3
                  i32.shl
                  i32.const 12
                  i32.add
                  local.get 6
                  i32.mul
                  i32.add
                  local.tee 3
                  i32.eqz
                  br_if 5 (;@2;)
                  local.get 6
                  i32.const 1
                  i32.add
                  local.set 6
                end
                local.get 3
                local.get 5
                i32.const 255
                i32.and
                local.tee 9
                i32.add
                i32.load8_u
                i32.eqz
                if  ;; label = @7
                  local.get 5
                  i32.const 1
                  i32.add
                  local.set 5
                  br 1 (;@6;)
                end
                local.get 13
                local.get 0
                i32.load offset=12
                local.tee 10
                local.get 9
                i32.mul
                local.get 3
                i32.add
                i32.const 12
                i32.add
                local.get 10
                call $memcpy
                local.set 8
                block  ;; label = @7
                  local.get 0
                  i32.load
                  local.get 7
                  i32.eq
                  if  ;; label = @8
                    local.get 11
                    local.get 9
                    local.get 0
                    i32.load offset=16
                    local.tee 9
                    i32.mul
                    local.get 10
                    i32.const 3
                    i32.shl
                    i32.add
                    local.get 3
                    i32.add
                    i32.const 12
                    i32.add
                    local.get 9
                    call $memcpy
                    drop
                    local.get 5
                    i32.const 1
                    i32.add
                    local.set 5
                    br 1 (;@7;)
                  end
                  local.get 0
                  i32.load offset=36
                  local.tee 9
                  i32.eqz
                  br_if 5 (;@2;)
                  local.get 5
                  i32.const 1
                  i32.add
                  local.set 5
                  local.get 0
                  local.get 8
                  local.get 11
                  local.get 8
                  local.get 10
                  local.get 0
                  i32.load offset=4
                  local.get 0
                  i32.load offset=32
                  local.get 9
                  call_indirect (type 0)
                  call $runtime.hashmapGet
                  i32.const 1
                  i32.and
                  i32.eqz
                  br_if 1 (;@6;)
                end
              end
              local.get 4
              i32.load offset=44
              local.tee 10
              i32.eqz
              br_if 3 (;@2;)
              local.get 4
              i32.const 8
              i32.add
              local.get 8
              local.get 11
              local.get 8
              local.get 4
              i32.load offset=20
              local.get 4
              i32.load offset=12
              local.get 4
              i32.load offset=40
              local.get 10
              call_indirect (type 0)
              call $runtime.hashmapSet
              br 1 (;@4;)
            end
          end
          local.get 0
          local.get 4
          i32.load offset=8
          i32.store
          local.get 0
          local.get 4
          i64.load offset=12 align=4
          i64.store offset=4 align=4
          local.get 0
          local.get 4
          i64.load offset=20 align=4
          i64.store offset=12 align=4
          local.get 0
          local.get 4
          i32.load8_u offset=28
          i32.store8 offset=20
          local.get 0
          local.get 4
          i64.load offset=32
          i64.store offset=24 align=4
          local.get 0
          local.get 4
          i64.load offset=40
          i64.store offset=32 align=4
          local.get 0
          i32.load offset=36
          local.tee 3
          i32.eqz
          br_if 1 (;@2;)
          local.get 1
          local.get 0
          i32.load offset=12
          local.get 0
          i32.load offset=4
          local.get 0
          i32.load offset=32
          local.get 3
          call_indirect (type 0)
          local.set 3
        end
        local.get 3
        i32.const 24
        i32.shr_u
        local.tee 5
        i32.const 1
        local.get 5
        select
        local.set 11
        i32.const 0
        local.set 6
        local.get 0
        local.get 3
        call $runtime.hashmapBucketAddrForHash
        local.set 3
        i32.const 0
        local.set 8
        i32.const 0
        local.set 10
        i32.const 0
        local.set 7
        loop  ;; label = @3
          local.get 3
          local.tee 5
          if  ;; label = @4
            i32.const 0
            local.set 3
            loop  ;; label = @5
              local.get 3
              i32.const 8
              i32.ne
              if  ;; label = @6
                local.get 0
                i32.load offset=12
                local.tee 6
                local.get 3
                i32.mul
                local.get 5
                i32.add
                i32.const 12
                i32.add
                local.set 12
                local.get 0
                i32.load offset=16
                local.get 3
                i32.mul
                local.get 6
                i32.const 3
                i32.shl
                i32.add
                local.get 5
                i32.add
                i32.const 12
                i32.add
                local.set 9
                block  ;; label = @7
                  local.get 3
                  local.get 5
                  i32.add
                  local.tee 13
                  i32.load8_u
                  local.tee 14
                  local.get 11
                  i32.ne
                  br_if 0 (;@7;)
                  local.get 0
                  i32.load offset=28
                  local.tee 15
                  i32.eqz
                  br_if 5 (;@2;)
                  local.get 1
                  local.get 12
                  local.get 6
                  local.get 0
                  i32.load offset=24
                  local.get 15
                  call_indirect (type 0)
                  i32.const 1
                  i32.and
                  i32.eqz
                  br_if 0 (;@7;)
                  local.get 9
                  local.get 2
                  local.get 0
                  i32.load offset=16
                  call $memcpy
                  drop
                  br 6 (;@1;)
                end
                local.get 7
                local.get 13
                local.get 8
                local.get 14
                i32.or
                local.tee 6
                select
                local.set 7
                local.get 8
                local.get 12
                local.get 6
                select
                local.set 8
                local.get 10
                local.get 9
                local.get 6
                select
                local.set 10
                local.get 3
                i32.const 1
                i32.add
                local.set 3
                br 1 (;@5;)
              end
            end
            local.get 5
            i32.load offset=8
            local.set 3
            local.get 5
            local.set 6
            br 1 (;@3;)
          end
        end
        local.get 8
        i32.eqz
        if  ;; label = @3
          local.get 0
          i32.load offset=16
          local.get 0
          i32.load offset=12
          i32.add
          i32.const 3
          i32.shl
          i32.const 12
          i32.add
          call $runtime.alloc
          local.set 3
          local.get 0
          local.get 0
          i32.load offset=8
          i32.const 1
          i32.add
          i32.store offset=8
          local.get 3
          i32.const 12
          i32.add
          local.get 1
          local.get 0
          i32.load offset=12
          local.tee 1
          call $memcpy
          local.get 1
          i32.const 3
          i32.shl
          i32.add
          local.get 2
          local.get 0
          i32.load offset=16
          call $memcpy
          drop
          local.get 3
          local.get 11
          i32.store8
          local.get 6
          i32.eqz
          br_if 1 (;@2;)
          local.get 6
          local.get 3
          i32.store offset=8
          br 2 (;@1;)
        end
        local.get 0
        local.get 0
        i32.load offset=8
        i32.const 1
        i32.add
        i32.store offset=8
        local.get 8
        local.get 1
        local.get 0
        i32.load offset=12
        call $memcpy
        drop
        local.get 10
        local.get 2
        local.get 0
        i32.load offset=16
        call $memcpy
        drop
        local.get 7
        i32.eqz
        br_if 0 (;@2;)
        local.get 7
        local.get 11
        i32.store8
        br 1 (;@1;)
      end
      call $runtime.nilPanic
      unreachable
    end
    local.get 4
    i32.const 48
    i32.add
    global.set $__stack_pointer)
  (func $runtime.runtimePanicAt (type 4) (param i32 i32)
    i32.const 65620
    i32.const 22
    call $runtime.printstring
    local.get 0
    local.get 1
    call $runtime.printstring
    i32.const 10
    call $runtime.putchar
    unreachable)
  (func $runtime.putchar (type 1) (param i32)
    (local i32 i32)
    i32.const 65764
    i32.load
    local.tee 1
    i32.const 119
    i32.le_u
    if  ;; label = @1
      i32.const 65764
      local.get 1
      i32.const 1
      i32.add
      local.tee 2
      i32.store
      local.get 1
      i32.const 65768
      i32.add
      local.get 0
      i32.store8
      local.get 0
      i32.const 255
      i32.and
      i32.const 10
      i32.ne
      i32.const 0
      local.get 1
      i32.const 119
      i32.ne
      select
      i32.eqz
      if  ;; label = @2
        i32.const 65716
        local.get 2
        i32.store
        i32.const 65764
        i32.const 0
        i32.store
      end
      return
    end
    i32.const 65665
    i32.const 18
    call $runtime.runtimePanicAt
    unreachable)
  (func $malloc.command_export (type 3) (param i32) (result i32)
    local.get 0
    call $malloc)
  (func $free.command_export (type 1) (param i32)
    local.get 0
    call $free)
  (func $calloc.command_export (type 2) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    i32.mul
    call $malloc)
  (func $realloc.command_export (type 2) (param i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32)
    block (result i32)  ;; label = @1
      local.get 0
      local.set 9
      global.get $__stack_pointer
      i32.const 16
      i32.sub
      local.tee 5
      global.set $__stack_pointer
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 1
            local.tee 7
            i32.eqz
            if  ;; label = @5
              local.get 9
              call $free
              br 1 (;@4;)
            end
            local.get 7
            i32.const 0
            i32.lt_s
            br_if 1 (;@3;)
            local.get 7
            call $runtime.alloc
            local.set 4
            local.get 9
            if  ;; label = @5
              local.get 5
              local.get 9
              i32.store offset=12
              local.get 5
              i32.const 12
              i32.add
              local.get 5
              call $runtime.hashmapBinaryGet
              i32.const 1
              i32.and
              i32.eqz
              br_if 3 (;@2;)
              local.get 5
              i32.load
              local.set 3
              block  ;; label = @6
                block  ;; label = @7
                  local.get 5
                  i32.load offset=4
                  local.tee 0
                  local.get 7
                  local.get 0
                  local.get 7
                  i32.lt_u
                  select
                  local.tee 0
                  i32.const 33
                  i32.lt_u
                  if  ;; label = @8
                    local.get 3
                    local.get 4
                    i32.eq
                    br_if 2 (;@6;)
                    local.get 3
                    local.get 0
                    local.get 4
                    i32.add
                    local.tee 2
                    i32.sub
                    i32.const 0
                    local.get 0
                    i32.const 1
                    i32.shl
                    i32.sub
                    i32.gt_u
                    br_if 1 (;@7;)
                  end
                  local.get 4
                  local.get 3
                  local.get 0
                  memory.copy
                  br 1 (;@6;)
                end
                local.get 3
                local.get 4
                i32.xor
                i32.const 3
                i32.and
                local.set 1
                block  ;; label = @7
                  block  ;; label = @8
                    local.get 3
                    local.get 4
                    i32.gt_u
                    if  ;; label = @9
                      local.get 1
                      if  ;; label = @10
                        local.get 0
                        local.set 1
                        local.get 4
                        local.set 2
                        br 3 (;@7;)
                      end
                      local.get 4
                      i32.const 3
                      i32.and
                      i32.eqz
                      if  ;; label = @10
                        local.get 0
                        local.set 1
                        local.get 4
                        local.set 2
                        br 2 (;@8;)
                      end
                      local.get 0
                      i32.eqz
                      br_if 3 (;@6;)
                      local.get 4
                      local.get 3
                      i32.load8_u
                      i32.store8
                      local.get 0
                      i32.const 1
                      i32.sub
                      local.set 1
                      local.get 4
                      i32.const 1
                      i32.add
                      local.tee 2
                      i32.const 3
                      i32.and
                      i32.eqz
                      if  ;; label = @10
                        local.get 3
                        i32.const 1
                        i32.add
                        local.set 3
                        br 2 (;@8;)
                      end
                      local.get 1
                      i32.eqz
                      br_if 3 (;@6;)
                      local.get 4
                      local.get 3
                      i32.load8_u offset=1
                      i32.store8 offset=1
                      local.get 0
                      i32.const 2
                      i32.sub
                      local.set 1
                      local.get 4
                      i32.const 2
                      i32.add
                      local.tee 2
                      i32.const 3
                      i32.and
                      i32.eqz
                      if  ;; label = @10
                        local.get 3
                        i32.const 2
                        i32.add
                        local.set 3
                        br 2 (;@8;)
                      end
                      local.get 1
                      i32.eqz
                      br_if 3 (;@6;)
                      local.get 4
                      local.get 3
                      i32.load8_u offset=2
                      i32.store8 offset=2
                      local.get 0
                      i32.const 3
                      i32.sub
                      local.set 1
                      local.get 4
                      i32.const 3
                      i32.add
                      local.tee 2
                      i32.const 3
                      i32.and
                      i32.eqz
                      if  ;; label = @10
                        local.get 3
                        i32.const 3
                        i32.add
                        local.set 3
                        br 2 (;@8;)
                      end
                      local.get 1
                      i32.eqz
                      br_if 3 (;@6;)
                      local.get 4
                      local.get 3
                      i32.load8_u offset=3
                      i32.store8 offset=3
                      local.get 4
                      i32.const 4
                      i32.add
                      local.set 2
                      local.get 3
                      i32.const 4
                      i32.add
                      local.set 3
                      local.get 0
                      i32.const 4
                      i32.sub
                      local.set 1
                      br 1 (;@8;)
                    end
                    block  ;; label = @9
                      local.get 1
                      br_if 0 (;@9;)
                      block  ;; label = @10
                        local.get 2
                        i32.const 3
                        i32.and
                        i32.eqz
                        br_if 0 (;@10;)
                        local.get 0
                        i32.eqz
                        br_if 4 (;@6;)
                        local.get 4
                        local.get 0
                        i32.const 1
                        i32.sub
                        local.tee 1
                        i32.add
                        local.tee 2
                        local.get 1
                        local.get 3
                        i32.add
                        i32.load8_u
                        i32.store8
                        local.get 2
                        i32.const 3
                        i32.and
                        i32.eqz
                        if  ;; label = @11
                          local.get 1
                          local.set 0
                          br 1 (;@10;)
                        end
                        local.get 1
                        i32.eqz
                        br_if 4 (;@6;)
                        local.get 4
                        local.get 0
                        i32.const 2
                        i32.sub
                        local.tee 1
                        i32.add
                        local.tee 2
                        local.get 1
                        local.get 3
                        i32.add
                        i32.load8_u
                        i32.store8
                        local.get 2
                        i32.const 3
                        i32.and
                        i32.eqz
                        if  ;; label = @11
                          local.get 1
                          local.set 0
                          br 1 (;@10;)
                        end
                        local.get 1
                        i32.eqz
                        br_if 4 (;@6;)
                        local.get 4
                        local.get 0
                        i32.const 3
                        i32.sub
                        local.tee 1
                        i32.add
                        local.tee 2
                        local.get 1
                        local.get 3
                        i32.add
                        i32.load8_u
                        i32.store8
                        local.get 2
                        i32.const 3
                        i32.and
                        i32.eqz
                        if  ;; label = @11
                          local.get 1
                          local.set 0
                          br 1 (;@10;)
                        end
                        local.get 1
                        i32.eqz
                        br_if 4 (;@6;)
                        local.get 4
                        local.get 0
                        i32.const 4
                        i32.sub
                        local.tee 0
                        i32.add
                        local.get 0
                        local.get 3
                        i32.add
                        i32.load8_u
                        i32.store8
                      end
                      local.get 0
                      i32.const 4
                      i32.lt_u
                      br_if 0 (;@9;)
                      local.get 0
                      i32.const 4
                      i32.sub
                      local.tee 1
                      i32.const 2
                      i32.shr_u
                      i32.const 1
                      i32.add
                      i32.const 3
                      i32.and
                      local.tee 2
                      if  ;; label = @10
                        local.get 3
                        i32.const 4
                        i32.sub
                        local.set 6
                        local.get 4
                        i32.const 4
                        i32.sub
                        local.set 8
                        loop  ;; label = @11
                          local.get 0
                          local.get 8
                          i32.add
                          local.get 0
                          local.get 6
                          i32.add
                          i32.load
                          i32.store
                          local.get 0
                          i32.const 4
                          i32.sub
                          local.set 0
                          local.get 2
                          i32.const 1
                          i32.sub
                          local.tee 2
                          br_if 0 (;@11;)
                        end
                      end
                      local.get 1
                      i32.const 12
                      i32.lt_u
                      br_if 0 (;@9;)
                      local.get 3
                      i32.const 16
                      i32.sub
                      local.set 6
                      local.get 4
                      i32.const 16
                      i32.sub
                      local.set 8
                      loop  ;; label = @10
                        local.get 0
                        local.get 8
                        i32.add
                        local.tee 1
                        i32.const 12
                        i32.add
                        local.get 0
                        local.get 6
                        i32.add
                        local.tee 2
                        i32.const 12
                        i32.add
                        i32.load
                        i32.store
                        local.get 1
                        i32.const 8
                        i32.add
                        local.get 2
                        i32.const 8
                        i32.add
                        i32.load
                        i32.store
                        local.get 1
                        i32.const 4
                        i32.add
                        local.get 2
                        i32.const 4
                        i32.add
                        i32.load
                        i32.store
                        local.get 1
                        local.get 2
                        i32.load
                        i32.store
                        local.get 0
                        i32.const 16
                        i32.sub
                        local.tee 0
                        i32.const 3
                        i32.gt_u
                        br_if 0 (;@10;)
                      end
                    end
                    local.get 0
                    i32.eqz
                    br_if 2 (;@6;)
                    local.get 0
                    local.tee 2
                    i32.const 3
                    i32.and
                    local.tee 6
                    if  ;; label = @9
                      local.get 3
                      i32.const 1
                      i32.sub
                      local.set 1
                      local.get 4
                      i32.const 1
                      i32.sub
                      local.set 8
                      loop  ;; label = @10
                        local.get 2
                        local.get 8
                        i32.add
                        local.get 1
                        local.get 2
                        i32.add
                        i32.load8_u
                        i32.store8
                        local.get 2
                        i32.const 1
                        i32.sub
                        local.set 2
                        local.get 6
                        i32.const 1
                        i32.sub
                        local.tee 6
                        br_if 0 (;@10;)
                      end
                    end
                    local.get 0
                    i32.const 4
                    i32.lt_u
                    br_if 2 (;@6;)
                    local.get 3
                    i32.const 4
                    i32.sub
                    local.set 3
                    local.get 4
                    i32.const 4
                    i32.sub
                    local.set 6
                    loop  ;; label = @9
                      local.get 2
                      local.get 6
                      i32.add
                      local.tee 0
                      i32.const 3
                      i32.add
                      local.get 2
                      local.get 3
                      i32.add
                      local.tee 1
                      i32.const 3
                      i32.add
                      i32.load8_u
                      i32.store8
                      local.get 0
                      i32.const 2
                      i32.add
                      local.get 1
                      i32.const 2
                      i32.add
                      i32.load8_u
                      i32.store8
                      local.get 0
                      i32.const 1
                      i32.add
                      local.get 1
                      i32.const 1
                      i32.add
                      i32.load8_u
                      i32.store8
                      local.get 0
                      local.get 1
                      i32.load8_u
                      i32.store8
                      local.get 2
                      i32.const 4
                      i32.sub
                      local.tee 2
                      br_if 0 (;@9;)
                    end
                    br 2 (;@6;)
                  end
                  local.get 1
                  i32.const 4
                  i32.lt_u
                  br_if 0 (;@7;)
                  local.get 1
                  i32.const 4
                  i32.sub
                  local.tee 6
                  i32.const 2
                  i32.shr_u
                  i32.const 1
                  i32.add
                  i32.const 7
                  i32.and
                  local.tee 0
                  if  ;; label = @8
                    loop  ;; label = @9
                      local.get 2
                      local.get 3
                      i32.load
                      i32.store
                      local.get 3
                      i32.const 4
                      i32.add
                      local.set 3
                      local.get 2
                      i32.const 4
                      i32.add
                      local.set 2
                      local.get 1
                      i32.const 4
                      i32.sub
                      local.set 1
                      local.get 0
                      i32.const 1
                      i32.sub
                      local.tee 0
                      br_if 0 (;@9;)
                    end
                  end
                  local.get 6
                  i32.const 28
                  i32.lt_u
                  br_if 0 (;@7;)
                  loop  ;; label = @8
                    local.get 2
                    local.get 3
                    i32.load
                    i32.store
                    local.get 2
                    local.get 3
                    i32.load offset=4
                    i32.store offset=4
                    local.get 2
                    local.get 3
                    i32.load offset=8
                    i32.store offset=8
                    local.get 2
                    local.get 3
                    i32.load offset=12
                    i32.store offset=12
                    local.get 2
                    local.get 3
                    i32.load offset=16
                    i32.store offset=16
                    local.get 2
                    local.get 3
                    i32.load offset=20
                    i32.store offset=20
                    local.get 2
                    local.get 3
                    i32.load offset=24
                    i32.store offset=24
                    local.get 2
                    local.get 3
                    i32.load offset=28
                    i32.store offset=28
                    local.get 3
                    i32.const 32
                    i32.add
                    local.set 3
                    local.get 2
                    i32.const 32
                    i32.add
                    local.set 2
                    local.get 1
                    i32.const 32
                    i32.sub
                    local.tee 1
                    i32.const 3
                    i32.gt_u
                    br_if 0 (;@8;)
                  end
                end
                local.get 1
                i32.eqz
                br_if 0 (;@6;)
                block  ;; label = @7
                  local.get 1
                  i32.const 7
                  i32.and
                  local.tee 6
                  i32.eqz
                  if  ;; label = @8
                    local.get 1
                    local.set 0
                    br 1 (;@7;)
                  end
                  local.get 1
                  local.set 0
                  loop  ;; label = @8
                    local.get 2
                    local.get 3
                    i32.load8_u
                    i32.store8
                    local.get 0
                    i32.const 1
                    i32.sub
                    local.set 0
                    local.get 2
                    i32.const 1
                    i32.add
                    local.set 2
                    local.get 3
                    i32.const 1
                    i32.add
                    local.set 3
                    local.get 6
                    i32.const 1
                    i32.sub
                    local.tee 6
                    br_if 0 (;@8;)
                  end
                end
                local.get 1
                i32.const 8
                i32.lt_u
                br_if 0 (;@6;)
                loop  ;; label = @7
                  local.get 2
                  local.get 3
                  i32.load8_u
                  i32.store8
                  local.get 2
                  local.get 3
                  i32.load8_u offset=1
                  i32.store8 offset=1
                  local.get 2
                  local.get 3
                  i32.load8_u offset=2
                  i32.store8 offset=2
                  local.get 2
                  local.get 3
                  i32.load8_u offset=3
                  i32.store8 offset=3
                  local.get 2
                  local.get 3
                  i32.load8_u offset=4
                  i32.store8 offset=4
                  local.get 2
                  local.get 3
                  i32.load8_u offset=5
                  i32.store8 offset=5
                  local.get 2
                  local.get 3
                  i32.load8_u offset=6
                  i32.store8 offset=6
                  local.get 2
                  local.get 3
                  i32.load8_u offset=7
                  i32.store8 offset=7
                  local.get 2
                  i32.const 8
                  i32.add
                  local.set 2
                  local.get 3
                  i32.const 8
                  i32.add
                  local.set 3
                  local.get 0
                  i32.const 8
                  i32.sub
                  local.tee 0
                  br_if 0 (;@7;)
                end
              end
              local.get 5
              local.get 9
              i32.store
              local.get 5
              call $runtime.hashmapBinaryDelete
            end
            local.get 5
            local.get 7
            i32.store offset=8
            local.get 5
            local.get 7
            i32.store offset=4
            local.get 5
            local.get 4
            i32.store
            local.get 5
            local.get 4
            i32.store offset=12
            local.get 5
            i32.const 12
            i32.add
            local.get 5
            call $runtime.hashmapBinarySet
          end
          local.get 5
          i32.const 16
          i32.add
          global.set $__stack_pointer
          local.get 4
          br 2 (;@1;)
        end
        call $runtime.slicePanic
        unreachable
      end
      i32.const 65592
      call $runtime._panic
      unreachable
    end)
  (func $_start.command_export (type 5)
    (local i32)
    memory.size
    local.set 0
    i32.const 65708
    i32.const 66192
    i32.store
    i32.const 65760
    local.get 0
    i32.const 16
    i32.shl
    i32.store
    i32.const 65760
    memory.size
    i32.const 16
    i32.shl
    i32.store
    call $user_entrypoint
    drop)
  (func $dummy.command_export (type 0) (param i32 i32 i32 i32) (result i32)
    i32.const 0)
  (func $user_entrypoint.command_export (type 3) (param i32) (result i32)
    call $user_entrypoint)
  (table (;0;) 3 3 funcref)
  (memory (;0;) 2)
  (global $__stack_pointer (mut i32) (i32.const 65536))
  (export "memory" (memory 0))
  (export "malloc" (func $malloc.command_export))
  (export "free" (func $free.command_export))
  (export "calloc" (func $calloc.command_export))
  (export "realloc" (func $realloc.command_export))
  (export "_start" (func $_start.command_export))
  (export "dummy" (func $dummy.command_export))
  (export "user_entrypoint" (func $user_entrypoint.command_export))
  (elem (;0;) (i32.const 1) func $runtime.memequal $runtime.hash32)
  (data $.rodata (i32.const 65536) "free: invalid pointer\00\00\00\00\00\01\00\15\00\00\00realloc: invalid pointer \00\01\00\18\00\00\00out of memorypanic: panic: runtime error: nil pointer dereferenceindex out of rangeslice out of range")
  (data $.data (i32.const 65704) "aQ\86A\90\02\01\00\e8\00\01\00\00\00\00\00t\01\01\00\c1\82\01\00\00\00\00\00\04\00\00\00\0c\00\00\00\01\00\00\00\00\00\00\00\01\00\00\00\00\00\00\00\02"))
